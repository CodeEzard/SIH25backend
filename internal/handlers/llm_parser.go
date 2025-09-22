package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"vericred/internal/models"
)

// ParseWithGemini uses Google's Gemini API to extract structured fields from raw OCR text.
func ParseWithGemini(ocrText string) (models.ParsedCredential, error) {
	var out models.ParsedCredential

	apiKey := os.Getenv("GEMINI_API_KEY")
	if strings.TrimSpace(apiKey) == "" {
		return out, errors.New("missing GEMINI_API_KEY")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return out, fmt.Errorf("failed to init Gemini client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash-lite")
	// Ask Gemini to return JSON only
	model.GenerationConfig = genai.GenerationConfig{ResponseMIMEType: "application/json"}

	prompt := `You are an expert data extraction assistant. Your job is to extract specific fields from the following raw text of an academic marksheet and return the data in a clean JSON format.

Here are the rules:
1. The required fields are: "register_number", "student_name", "course_name", "year_of_passing", and "university_name".
2. If a field cannot be found in the text, its value in the JSON must be null.
3. Your entire response must be ONLY the JSON object. Do not include any explanations, apologies, or any text before or after the JSON.
4. Clean the extracted data by removing any unnecessary newline characters or extra whitespace.

Here is the raw text:
"""
[INSERT RAW OCR TEXT HERE]
"""`

	prompt = strings.Replace(prompt, "[INSERT RAW OCR TEXT HERE]", ocrText, 1)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return out, fmt.Errorf("gemini generation failed: %w", err)
	}
	if resp == nil || len(resp.Candidates) == 0 || resp.Candidates[0] == nil || resp.Candidates[0].Content == nil {
		return out, errors.New("empty response from Gemini")
	}

	var sb strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		if t, ok := part.(genai.Text); ok {
			sb.WriteString(string(t))
		} else {
			sb.WriteString(fmt.Sprint(part))
		}
	}
	jsonStr := strings.TrimSpace(sb.String())
	if jsonStr == "" {
		return out, errors.New("no text in Gemini response")
	}

	// Normalize: strip code fences and extract the first JSON object if needed
	jsonStr = stripCodeFences(jsonStr)
	if candidate, ok := extractFirstJSON(jsonStr); ok {
		jsonStr = candidate
	}

	// Tolerate nulls by unmarshaling into a map[string]any first
	var tmp map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &tmp); err != nil {
		return out, fmt.Errorf("failed to parse Gemini JSON: %w", err)
	}
	get := func(k string) string {
		v, ok := tmp[k]
		if !ok || v == nil {
			return ""
		}
		switch t := v.(type) {
		case string:
			return strings.TrimSpace(t)
		default:
			b, _ := json.Marshal(t)
			return strings.TrimSpace(string(b))
		}
	}

	out.RegisterNumber = get("register_number")
	out.StudentName = get("student_name")
	out.CourseName = get("course_name")
	out.YearOfPassing = get("year_of_passing")
	out.UniversityName = get("university_name")

	if strings.TrimSpace(out.RegisterNumber) == "" {
		return out, errors.New("register number not found")
	}
	return out, nil
}

// stripCodeFences removes surrounding Markdown code fences like ```json ... ```.
func stripCodeFences(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "```") {
		// drop leading backticks and optional language tag
		s = strings.TrimPrefix(s, "```")
		s = strings.TrimSpace(s)
		// remove a possible language tag at the start of the fence
		if i := strings.IndexByte(s, '\n'); i != -1 {
			first := strings.TrimSpace(s[:i])
			if len(first) > 0 && len(first) < 20 { // likely a language tag like json
				s = s[i+1:]
			}
		}
		// strip trailing fence if present
		if idx := strings.LastIndex(s, "```"); idx != -1 {
			s = s[:idx]
		}
	}
	return strings.TrimSpace(s)
}

// extractFirstJSON attempts to extract the first balanced JSON object or array.
func extractFirstJSON(s string) (string, bool) {
	if obj, ok := extractBalanced(s, '{', '}'); ok {
		return obj, true
	}
	if arr, ok := extractBalanced(s, '[', ']'); ok {
		return arr, true
	}
	return "", false
}

func extractBalanced(s string, open, close rune) (string, bool) {
	start := -1
	depth := 0
	for i, r := range s {
		if r == open {
			if depth == 0 {
				start = i
			}
			depth++
		} else if r == close {
			if depth > 0 {
				depth--
				if depth == 0 && start != -1 {
					return s[start : i+1], true
				}
			}
		}
	}
	return "", false
}

package models

// ParsedCredential holds the normalized fields parsed from OCR text
// for subsequent verification against the database.
type ParsedCredential struct {
	RegisterNumber string `json:"register_number"`
	StudentName    string `json:"student_name"`
	CourseName     string `json:"course_name"`
	YearOfPassing  string `json:"year_of_passing"`
	UniversityName string `json:"university_name"`
}

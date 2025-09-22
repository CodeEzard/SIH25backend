package googlevision

import (
	"context"
	"fmt"
	"log"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
)

func ImgOcr() {
	image_url := "https://imgv2-1-f.scribdassets.com/img/document/701163089/original/901511559d/1?v=1"

	ctx := context.Background()
	log.Printf("Step 1: Fetching image from URL: %s", image_url)
	
	response, err := http.Get(image_url)	
	if err != nil {
		log.Fatalf("FATAL: Failed to fetch image from URL: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("FATAL: Received non-200 status code: %d", response.StatusCode)
	}

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("FATAL: Failed to create Vision API client: %v", err)
	}
	defer client.Close()

	image, err := vision.NewImageFromReader(response.Body)
	if err != nil {
		log.Fatalf("FATAL: Failed to create image object from response: %v", err)
	}

	annotation, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		log.Fatalf("FATAL: Vision API failed to detect text: %v", err)
	}

	log.Println("Step 3: Vision API call successful. Printing result.")

	if annotation == nil {
		fmt.Println("No text was found in the image.")
	} else {
		fmt.Println("\n================ OCR RESULT ================")
		fmt.Println(annotation.Text)
		fmt.Println("==========================================")
	}
}
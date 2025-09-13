package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file upload function called.")

	r.ParseMultipartForm(10 << 20) // dont upload shit greater than 10 mb mf.

	file, handler, err := r.FormFile("randomfile")
	if err != nil {
		fmt.Println("Error retrieving the file.")
		fmt.Println(err)
		return
	}

	defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := os.CreateTemp("files", "*")
	if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tempFile.Write(fileBytes)
	fmt.Fprint(w, "successfully uploaded file")

}
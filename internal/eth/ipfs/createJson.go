package ipfs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func CreateJSONFileAndStoreToIPFS(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "vericred-*.json")
	if err != nil {
		http.Error(w, "failed to create temp file", http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()

	if _, err = tmpFile.Write(bodyBytes); err != nil {
		http.Error(w, "failed to write temp file", http.StatusInternalServerError)
		return
	}

	msg, err := UploadToIPFS(tmpFile.Name())
	if err != nil {
		http.Error(w, "failed to upload to IPFS", http.StatusInternalServerError)
		return
	}

	fmt.Println(msg)
	json.NewEncoder(w).Encode(map[string]string{"ipfslink": msg})
}
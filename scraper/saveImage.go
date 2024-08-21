package fragranitca

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SaveImage(url, path string) {
	dir := "downloads"

	splittedUrl := strings.Split(url, "/")
	fileName := splittedUrl[len(splittedUrl)-1]

	filePath := filepath.Join(dir, fileName)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Image saved to %s", filePath)
}

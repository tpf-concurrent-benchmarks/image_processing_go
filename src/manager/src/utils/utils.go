package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetImagesInDirectory(directory string) []string {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		log.Fatalf("Directory %s does not exist or is not a directory", directory)
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var imageFiles []string
	for _, file := range files {
		if !file.IsDir() {
			extension := strings.ToLower(filepath.Ext(file.Name()))
			if extension == ".jpg" || extension == ".png" || extension == ".gif" {
				imageFiles = append(imageFiles, filepath.Join(directory, file.Name()))
			}
		}
	}

	return imageFiles
}

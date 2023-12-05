package image_processing

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func ChangeResolution(inputPath, outputPath string, width, height int) {

	jpegImage, err := imaging.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening image for resolutio change:", err)
		return
	}

	resizedImage := imaging.Resize(jpegImage, width, height, imaging.Linear)

	err = imaging.Save(resizedImage, outputPath)
	if err != nil {
		fmt.Println("Error saving image with changed resolution:", err)
		return
	}

}

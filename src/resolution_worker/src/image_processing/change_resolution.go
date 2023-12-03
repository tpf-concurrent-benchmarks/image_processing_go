package image_processing

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func ChangeResolution(input_path, output_path string, width, height int) {
	// Open the JPEG image
	jpegImage, err := imaging.Open(input_path)
	if err != nil {
		fmt.Println("Error opening JPEG image:", err)
		return
	}

	// Resize the image
	resizedImage := imaging.Resize(jpegImage, width, height, imaging.Linear)

	// Save the image as PNG
	err = imaging.Save(resizedImage, output_path)
	if err != nil {
		fmt.Println("Error saving PNG image:", err)
		return
	}

}
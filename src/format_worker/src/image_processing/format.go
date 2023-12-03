package image_processing

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func Format(input_path, output_path string) {
	// Open the JPEG image
	jpegImage, err := imaging.Open(input_path)
	if err != nil {
		fmt.Println("Error opening JPEG image:", err)
		return
	}

	// Save the image as PNG
	err = imaging.Save(jpegImage, output_path)
	if err != nil {
		fmt.Println("Error saving PNG image:", err)
		return
	}

}
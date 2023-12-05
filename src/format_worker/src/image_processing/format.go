package image_processing

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func Format(inputPath, outputPath string) {

	jpegImage, err := imaging.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening JPEG image:", err)
		return
	}

	err = imaging.Save(jpegImage, outputPath)
	if err != nil {
		fmt.Println("Error saving PNG image:", err)
		return
	}

}

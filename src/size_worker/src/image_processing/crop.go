package image_processing

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func CropCentered(inputPath, outputPath string, width, height int) {

	pngImage, err := imaging.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening image for crop:", err)
		return
	}

	croppedImage := imaging.CropAnchor(pngImage, width, height, imaging.Center)

	err = imaging.Save(croppedImage, outputPath)
	if err != nil {
		fmt.Println("Error saving cropped image:", err)
		return
	}

}

package image_processing

import (
	"fmt"
	"github.com/disintegration/imaging"
)

func CropCentered(input_path, output_path string, width, height int) {

	pngImage, err := imaging.Open(input_path)
	if err != nil {
		fmt.Println("Error opening image for crop:", err)
		return
	}

	croppedImage := imaging.CropAnchor(pngImage, width, height, imaging.Center)

	err = imaging.Save(croppedImage, output_path)
	if err != nil {
		fmt.Println("Error saving croped image:", err)
		return
	}

}
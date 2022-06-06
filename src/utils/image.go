package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// Writes image to disk as a jpeg
// Do not include the extension in the filename
func SaveImageToDisk(fName string, img *image.Image) error {
	// save img to disk
	f, err := os.Create(fmt.Sprintf("%s.jpg", fName))
	if err != nil {
		return err
	}

	if err = jpeg.Encode(f, *img, nil); err != nil {
		return err
	}

	f.Close()
	return nil
}

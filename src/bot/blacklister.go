package bot

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/CarlFlo/blacklisterBot/src/database"
	"github.com/CarlFlo/blacklisterBot/src/utils"
	"github.com/CarlFlo/malm"
	"github.com/corona10/goimagehash"
)

func checkImage(img *image.Image) bool {

	var found bool
	var err error

	// Check the SHA-1 first

	found, err = sha1Check(img)
	if err != nil {
		malm.Error("%s", err)
	} else if found {
		return true
	}

	// check Average, Difference & Perception
	found, err = averageDifferencePerceptionCheck(img)
	if err != nil {
		malm.Error("%s", err)
	} else if found {
		return true
	}

	return false
}

func sha1Check(img *image.Image) (bool, error) {

	// convert *image.Image to a []byte
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, *img, nil)
	if err != nil {
		return false, err
	}

	hasher := sha1.New()
	hasher.Write(buf.Bytes())
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	// Check database
	return database.SearchSHA1(hash)
}

func averageDifferencePerceptionCheck(img *image.Image) (bool, error) {

	aHash, err := goimagehash.AverageHash(*img)
	if err != nil {
		return false, err
	}

	dHash, err := goimagehash.DifferenceHash(*img)
	if err != nil {
		return false, err
	}

	pHash, err := goimagehash.PerceptionHash(*img)
	if err != nil {
		return false, err
	}

	found, err := database.SearchAveragePerceptionDifference(aHash, dHash, pHash)
	if err != nil {
		return false, err
	}

	return found, nil
}

func handleImage(url *string) (*image.Image, error) {

	// converty to []byte to image.Image
	b, err := utils.FetchFromURL(url)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return &img, nil
}

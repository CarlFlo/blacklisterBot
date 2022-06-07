package bot

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"

	"github.com/CarlFlo/blacklisterBot/src/database"
	"github.com/CarlFlo/malm"
	"github.com/corona10/goimagehash"
)

func checkImage(img *image.Image) bool {

	// Check the SHA-1 first

	found, err := sha1Check(img)
	malm.Info("SHA-1: %v", found)
	if err != nil {
		malm.Error("%s", err)
	} else if found {
		return true
	}

	// check average

	// check difference

	// check perception

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

func averageCheck(img *image.Image) (bool, error) {

	averageHash, err := goimagehash.AverageHash(*img)
	if err != nil {
		return false, err
	}

	found := database.SearchAverage(averageHash)

	return found, nil
}

func handleImage(url *string) (*image.Image, error) {

	// converty to []byte to image.Image
	b, err := getAttatchment(url)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return &img, nil
}

func getAttatchment(url *string) ([]byte, error) {

	resp, err := http.Get(*url)
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

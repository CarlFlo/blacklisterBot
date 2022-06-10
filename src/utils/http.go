package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchFromURL(url *string) ([]byte, error) {

	resp, err := http.Get(*url)
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

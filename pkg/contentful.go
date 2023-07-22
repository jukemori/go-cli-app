// pkg/contentful.go

package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// FetchContentfulData fetches data from the Contentful API and returns the response as a string.
func FetchContentfulData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch data from Contentful: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Contentful API returned status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}

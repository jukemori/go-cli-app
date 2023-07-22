package pkg

// pkg/contentful.go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ContentfulResponse represents the JSON response from the Contentful API.
type ContentfulResponse struct {
	Sys struct {
		ID        string `json:"id"`
		CreatedAt string `json:"createdAt"`
	} `json:"sys"`
	Fields struct {
		Name string `json:"name"`
	} `json:"fields"`
}

// FetchContentfulData fetches data from the Contentful API and returns the response as a ContentfulResponse struct.
func FetchContentfulData(url string) (*ContentfulResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from Contentful: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Contentful API returned status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var contentfulResponse ContentfulResponse
	err = json.Unmarshal(body, &contentfulResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing Contentful response: %v", err)
	}

	return &contentfulResponse, nil
}

package src2struct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Options struct {
	Client             *http.Client
	Header             *http.Header
	IncludeBodyInError bool
}

// LoadFileYaml will load a Yaml file and unmarshal it into the provided interface
func LoadUrlJson(logger *log.Entry, options *Options, url string, data interface{}) error {
	// Create/set the http client
	var httpClient *http.Client
	if options == nil || options.Client == nil {
		httpClient = &http.Client{}
	} else {
		httpClient = options.Client
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("when creating a http client based on the url '%s': %w", url, err)
	}

	// Add headers (if any)
	if options != nil && options.Header != nil {
		for key, values := range *options.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	// Execute the request, retreive the result (body) and parse it into the received object
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("when calling url '%s': %w", url, err)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return fmt.Errorf("when reading the body of the response from url '%s': %w", url, err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("got status code %d and body '%s' when calling url '%s'", resp.StatusCode, string(body), url)
	}
	if err := json.Unmarshal(body, data); err != nil {
		if options != nil && options.IncludeBodyInError {
			return fmt.Errorf("when converting body '%s' to a json object: %w", string(body), err)
		} else {
			return fmt.Errorf("when converting the returned body to a json object: %w", err)
		}
	}
	return nil
}

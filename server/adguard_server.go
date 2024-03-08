package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

// HTTPTimeout timeout to call endpoints
const HTTPTimeout = 2

const urlTemplate = "http://%s:%d/%s"

// AdguardServer plex server information
type AdguardServer struct {
	Host       string
	Port       int
	Token      string
	HTTPClient http.Client
}

// cache responses for 5 seconds
var cacheResponse = cache.New(5*time.Second, 10*time.Second)

// SendRequest send requests to plex endpoints
func (ad AdguardServer) SendRequest(api string, jsonStruct interface{}) error {

	url := fmt.Sprintf(urlTemplate, ad.Host, ad.Port, api)

	var reponseBody []byte

	cachedResponse, found := cacheResponse.Get(url)

	if found {
		reponseBody = cachedResponse.([]byte)
		log.Printf("URL: %s hit cache", url)
	} else {
		log.Printf("URL: %s miss cache", url)
		request, _ := http.NewRequest(http.MethodGet, url, nil)

		if len(strings.TrimSpace(ad.Token)) != 0 {
			request.Header.Add("Authorization", fmt.Sprintf("Basic %s", ad.Token))
		}

		request.Header.Add("Accept", "application/json")
		response, err := ad.HTTPClient.Do(request)
		if err != nil {
			return err
		}

		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("error: status code %d from server", response.StatusCode)
		}

		reponseBody, err = io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		cacheResponse.Set(url, reponseBody, 5*time.Second)

		defer response.Body.Close()
	}

	if err := json.Unmarshal([]byte(reponseBody), &jsonStruct); err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}
	return nil

}

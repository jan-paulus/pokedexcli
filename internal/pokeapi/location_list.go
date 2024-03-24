package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocations(pageURL *string) (LocationListResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		cachedData := LocationListResponse{}
		err := json.Unmarshal(val, &cachedData)
		if err != nil {
			return LocationListResponse{}, err
		}

		return cachedData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationListResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationListResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return LocationListResponse{}, fmt.Errorf("Response failed with status code %d and\nbody: %s\n", res.StatusCode, data)
	}
	if err != nil {
		return LocationListResponse{}, err
	}

	resData := LocationListResponse{}
	err = json.Unmarshal(data, &resData)

	if err != nil {
		return LocationListResponse{}, err
	}

	c.cache.Add(url, data)
	return resData, nil
}

type LocationListResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

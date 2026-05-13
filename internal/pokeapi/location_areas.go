package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	url := baseUrl + "/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}
	if val, ok := c.cache.Get(url); ok {
		cacheJson := LocationAreaResponse{}
		err := json.Unmarshal(val, &cacheJson)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return cacheJson, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return LocationAreaResponse{}, err
	}
	locationAreas := LocationAreaResponse{}
	errJson := json.Unmarshal(body, &locationAreas)
	if errJson != nil {
		return LocationAreaResponse{}, nil
	}
	c.cache.Add(url, body)
	return locationAreas, nil
}

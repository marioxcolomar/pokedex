package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemonPerLocationArea(locationId string) (ExploreResponse, error) {
	url := baseUrl + "/location-area/" + locationId

	if val, ok := c.cache.Get(url); ok {
		cache := ExploreResponse{}
		err := json.Unmarshal(val, &cache)
		if err != nil {
			return ExploreResponse{}, err
		}
		return cache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreResponse{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreResponse{}, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return ExploreResponse{}, err
	}
	exploreResponse := ExploreResponse{}
	errJson := json.Unmarshal(body, &exploreResponse)
	if errJson != nil {
		return ExploreResponse{}, nil
	}

	c.cache.Add(url, body)

	return exploreResponse, nil
}

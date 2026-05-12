package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pokedex/internal/pokecache"
)

func (c *Client) ListLocationAreas(pageUrl *string) (pokecache.LocationAreaResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	if val, ok := c.cache.Get(url); ok {
		cacheJson := pokecache.LocationAreaResponse{}
		err := json.Unmarshal(val, &cacheJson)
		if err != nil {
			return pokecache.LocationAreaResponse{}, err
		}
		return cacheJson, nil
	}

	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokecache.LocationAreaResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.Response.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.Response.StatusCode, body)
		return pokecache.LocationAreaResponse{}, err
	}
	if err != nil {
		log.Fatal(err)
		return pokecache.LocationAreaResponse{}, err
	}
	locationAreas := pokecache.LocationAreaResponse{}
	errJson := json.Unmarshal(body, &locationAreas)
	if errJson != nil {
		return pokecache.LocationAreaResponse{}, nil
	}
	c.cache.Add(url, body)
	return locationAreas, nil
}

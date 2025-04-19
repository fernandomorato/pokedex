package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaPage struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreaPage(pageURL *string) (LocationAreaPage, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if cachedData, ok := c.cache.Get(url); ok {
		result := LocationAreaPage{}
		err := json.Unmarshal(cachedData, &result)
		if err != nil {
			return LocationAreaPage{}, nil
		}
		return result, nil
	}

	// Build get request for location page
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaPage{}, err
	}

	// Get response from request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaPage{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaPage{}, fmt.Errorf("error: %s", res.Status)
	}

	// Unmarshal location page
	rawPage, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaPage{}, err
	}
	page := LocationAreaPage{}
	err = json.Unmarshal(rawPage, &page)
	if err != nil {
		return LocationAreaPage{}, err
	}
	c.cache.Add(url, rawPage)
	return page, nil
}

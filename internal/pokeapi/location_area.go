package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// LocationAreas sometimes Next and Previous could be null, thus using pointer for these two fields
type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint + "/" + locationAreaName

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Printf("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			panic("error closing response body")
		}
	}()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullURL, dat)
	return locationArea, nil
}

func (c *Client) LocationAreaList(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		locationAreas := LocationAreas{}
		err := json.Unmarshal(dat, &locationAreas)
		if err != nil {
			return LocationAreas{}, err
		}
		return locationAreas, nil
	}
	fmt.Printf("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			panic("error closing response body")
		}
	}()

	if res.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(dat, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}
	c.cache.Add(fullURL, dat)
	return locationAreas, nil
}

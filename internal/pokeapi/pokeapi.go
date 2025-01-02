package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Drnel/btdv_go_pokedex/internal/pokecache"
)

type Location_area_list struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func PrintLANames(url string, cache *pokecache.Cache) (previous string, next string) {
	data, ok := cache.Get(url)
	if !ok {

		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		data = body
		fmt.Println("Had to use the internet 🌐")
	}
	cache.Add(url, data)
	Location_areas := Location_area_list{}
	err := json.Unmarshal(data, &Location_areas)
	if err != nil {
		fmt.Println(err)
	}

	for _, location_area := range Location_areas.Results {
		fmt.Println(location_area.Name)
	}
	return Location_areas.Previous, Location_areas.Next
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func printNames(url string) (previous string, next string) {
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

	Location_areas := Location_area_list{}
	err = json.Unmarshal(body, &Location_areas)
	if err != nil {
		fmt.Println(err)
	}

	for _, location_area := range Location_areas.Results {
		fmt.Println(location_area.Name)
	}
	return Location_areas.Previous, Location_areas.Next
}

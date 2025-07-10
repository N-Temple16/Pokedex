package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config, strings []string) error {
	var body []byte
	var err error

	data, ok := cache.Get(config.Next)
	if ok {
		body = data
	} else {
		res, err := http.Get(config.Next)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(config.Next, bodyBytes)
		body = bodyBytes
	}

	locations := LocationResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	config.Next = locations.Next
	if locations.Previous != nil {
		config.Previous = *locations.Previous
	} else {
		config.Previous = ""
	}	

    return nil
}
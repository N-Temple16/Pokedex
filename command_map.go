package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	res, err := http.Get(config.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
        return err
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
	fmt.Println("Next URL:", config.Next)
	if locations.Previous != nil {
		config.Previous = *locations.Previous
	} else {
		config.Previous = ""
	}	

    return nil
}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(config *Config) error {
	if config.Previous == "" {
        fmt.Println("you're on the first page")
        return nil
    }
	
	res, err := http.Get(config.Previous)
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
	if locations.Previous != nil {
		config.Previous = *locations.Previous
	} else {
		config.Previous = ""
	}	

    return nil
}
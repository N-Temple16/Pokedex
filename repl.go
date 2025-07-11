package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)

    config := &Config{
            Next:     "https://pokeapi.co/api/v2/location-area/",
            Previous: "",
        }

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()

        words := cleanInput(scanner.Text())
        if len(words) == 0 {
            continue
        }

        commandName := words[0]
        location := words[1:]

        command, exists := getCommands()[commandName]

        if exists {
            err := command.callback(config, location)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
		    fmt.Print("Unkown command")
            continue
	    }
    }
}

func cleanInput(text string) []string {
    var stringSlice []string
    removeSpace := strings.Trim(text, " ")
    textLower := strings.ToLower(removeSpace)
    words := strings.Fields(textLower)
    for _, word := range words {
        stringSlice = append(stringSlice, word)
    }
    return stringSlice
}

type cliCommand struct {
    name        string
    description string
    callback    func(config *Config, strings []string) error
}

type Config struct {
    Next     string
    Previous string
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type SpecifiedLocation struct {
	Encounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
    BaseExperience int `json:"base_experience"`
	Name   string `json:"name"`
    Height int `json:"height"`
    Weight int `json:"weight"`
    Stats []PokemonStats `json:"stats"`
    Types []PokemonTypes `json:"types"`
}

type PokemonStats struct {
    BaseStat int `json:"base_stat"`
    StatName StatNames `json:"stat"`
}

type StatNames struct {
    Name string `json:"name"`
}

type PokemonTypes struct {
    TypeName TypeNames `json:"type"`
}

type TypeNames struct {
    Name string `json:"name"`
}


func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "map": {
            name:        "map",
            description: "Displays 20 upcoming locations",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays 20 previous locations",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "Displays list of all Pokemon located in an area",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch",
            description: "Catch and add a Pokemon to the Pokedex",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect",
            description: "Inspect a caught Pokemon's information",
            callback:    commandInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "Displays all caught Pokemon",
            callback:    commandPokedex,
        },
    }   
}
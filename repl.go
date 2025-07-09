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

        command, exists := getCommands()[commandName]

        if exists {
            err := command.callback(config)
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
    callback    func(config *Config) error
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
    }   
}
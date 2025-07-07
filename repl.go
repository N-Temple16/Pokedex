package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)

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
            err := command.callback()
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
    callback    func() error
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
    }   
}
package main

import (
	"fmt"
)

func commandHelp(config *Config, strings []string) error {
	fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, command := range getCommands() {
       fmt.Printf("%s: %s\n", command.name, command.description) 
    }
	fmt.Println()
    return nil
}
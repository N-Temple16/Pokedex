package main

import (
	"strings"
)

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
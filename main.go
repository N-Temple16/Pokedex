package main

import (
    "github.com/Nigel-Temple16/pokedex/internal/pokecache"
    "time"
)

var cache *pokecache.Cache
var caughtPokemon map[string]Pokemon

func main() {
    cache = pokecache.NewCache(time.Second * 5)
    caughtPokemon = make(map[string]Pokemon)
    startRepl()
}

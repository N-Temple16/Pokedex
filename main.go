package main

import (
    "github.com/Nigel-Temple16/pokedex/internal/pokecache"
    "time"
)

var cache *pokecache.Cache

func main() {
    cache = pokecache.NewCache(time.Second * 5)
    startRepl()
}

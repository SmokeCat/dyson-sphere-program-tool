package main

import (
	"dyson-sphere-program-tool/internal/gamedata"
	"log"
)

func main() {
	data, err := gamedata.LoadGameData()
	if err != nil {
		log.Fatal(err)
	}
	println(data.Version)
}

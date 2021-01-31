package gamedata

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Data struct {
	Version      string         `json:"version"`
	Material     []Material     `json:"material"`
	Component    []Component    `json:"component"`
	Architecture []Architecture `json:"architecture"`
	Blueprint    []Blueprint    `json:"blueprint"`
}

func LoadGameData() (*Data, error) {
	data := Data{}
	content, err := ioutil.ReadFile("configs/game_data.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal(err)
	}
	return &data, err
}

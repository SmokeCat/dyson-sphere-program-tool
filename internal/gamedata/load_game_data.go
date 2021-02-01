package gamedata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Data struct {
	Version      string         `json:"version"`
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

// func (d *Data) GetBlueprintById(id int) (Blueprint, error) {
// 	for i:= 0; i < len(d.Blueprint); i++ {
// 		if d.Blueprint[i].Id == id {
// 			return d.Blueprint[i], nil
// 		}
// 	}
// 	log.Fatal("can not find blueprint: id ", id)
// 	return Blueprint{}, errors.New("can not find blueprint: id")
// }

func (d *Data) PrintCost(bpId int) {
	fmt.Println("print cost")
}

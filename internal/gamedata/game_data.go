package gamedata

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Component 组件
type Component struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Architecture 建筑
type Architecture struct {
	Name string `json:"name"`
}

// Blueprint 合成蓝图
type Blueprint struct {
	Time    float32        `json:"time"`
	Product map[string]int `json:"product"`
	Cost    map[string]int `json:"cost"`
}

// GameData 游戏数据，映射 game_data.json 文件，用于维护游戏数据
type GameData struct {
	Version      string         `json:"version"`
	Component    []Component    `json:"component"`
	Architecture []Architecture `json:"architecture"`
	Blueprint    []Blueprint    `json:"blueprint"`
}

// LoadGameData 加载游戏数据
func LoadGameData() (*GameData, error) {
	data := GameData{}

	content, err := ioutil.ReadFile("configs/game_data.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	return &data, err
}

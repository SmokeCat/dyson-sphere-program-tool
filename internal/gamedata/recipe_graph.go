package gamedata

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// RecipeGraphItem 配方图节点
type RecipeGraphItem struct {
	ID      int                      `json:"id"`
	Product []map[string]interface{} `json:"product"`
	Cost    []map[string]interface{} `json:"cost"`
}

// RecipeGraph 配方图，由 GameData 生成，用于本工具数据计算
type RecipeGraph = map[string]*RecipeGraphItem

// Recipe 全局RecipeGraph对象
var Recipe *RecipeGraph

// LoadRecipeGraph 加载配方图
func LoadRecipeGraph() (*RecipeGraph, error) {
	Recipe = new(RecipeGraph)

	content, err := ioutil.ReadFile("configs/recipe_graph.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &Recipe)
	if err != nil {
		log.Fatal(err)
	}

	return Recipe, err
}

// InitRecipeGraph 根据 game_data.json 生成配方图
func InitRecipeGraph() {
	data, err := LoadGameData()
	if err != nil {
		log.Fatal(err)
	}

	components := data.Component
	architectures := data.Architecture
	blueprints := data.Blueprint

	Recipe := make(RecipeGraph)
	var id int
	// 1001 - 1999 组件
	for id, i, j := 1001, 0, len(components); i < j; i, id = i+1, id+1 {
		name := components[i].Name
		gen := RecipeGraphItem{
			ID:      id,
			Product: make([]map[string]interface{}, 0),
			Cost:    make([]map[string]interface{}, 0),
		}
		Recipe[name] = &gen
	}
	if id > 1999 {
		log.Fatal("component id greater than 1999 !")
	}

	// 2001 - 2999 建筑
	for id, i, j := 2001, 0, len(architectures); i < j; i, id = i+1, id+1 {
		name := architectures[i].Name
		gen := RecipeGraphItem{
			ID:      id,
			Product: make([]map[string]interface{}, 0),
			Cost:    make([]map[string]interface{}, 0),
		}
		Recipe[name] = &gen
	}
	if id > 2999 {
		log.Fatal("architecture id greater than 2999!")
	}

	// 3001 - 3999 合成配方
	for id, i, j := 3001, 0, len(blueprints); i < j; i, id = i+1, id+1 {
		time := blueprints[i].Time
		p := blueprints[i].Product
		c := blueprints[i].Cost

		// 遍历 product
		for k, v := range p {

			item := make(map[string]interface{})
			item["time"] = time
			item["count"] = v

			// 添加生产材料
			for ik, iv := range c {
				item[ik] = iv
			}

			Recipe[k].Product = append(Recipe[k].Product, item)
		}
	}

	// 写入文件
	res, err := json.MarshalIndent(&Recipe, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("configs/recipe_graph.json", res, 0644); err != nil {
		log.Fatal(err)
	}
}

// RecurCost 递归获取消耗
func (r *RecipeGraphItem) RecurCost() []map[string]float64 {
	record := make(map[string]float64)
	res := make([]map[string]float64, 0)
	res = append(res, record)

	recurCost(r, &res, &record)

	return res
}

// recurCost 递归查询消耗，私有方法
func recurCost(item *RecipeGraphItem, records *[]map[string]float64, record *map[string]float64) {
	products := (*item).Product
	pLen := len(products)
	if pLen < 1 {
		return
	}
	for _, v := range products {
		for it, count := range v {
			if it != "time" && it != "count" {
				(*record)[it] += count.(float64) / v["time"].(float64)
				recurCost((*Recipe)[it], records, record)
			}
		}
	}
}

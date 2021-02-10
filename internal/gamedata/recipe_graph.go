package gamedata

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/smokecat/dyson-sphere-program-tool/internal/util"
)

// RecipeGraph 配方图，由 GameData 生成，用于本工具数据计算
type RecipeGraph = map[string]*RecipeGraphItem

// RecipeGraphItem 配方图节点
type RecipeGraphItem struct {
	ID      int
	Product []recipeGraphProductItem
	Cost    []interface{}
}

type recipeGraphProductItem struct {
	Count int
	Time  float64
	Items []recipeGraphItemWithCount
}

type recipeGraphItemWithCount struct {
	Count int
	Name  string
	Item  *RecipeGraphItem
}

// Recipe 全局RecipeGraph对象
var Recipe *RecipeGraph

// LoadRecipeGraph 加载配方图
func LoadRecipeGraph() (*RecipeGraph, error) {
	Recipe = &RecipeGraph{}

	// 读取配方图JSON文件
	recipeGraphJson := make(map[string]map[string]interface{})
	content, err := ioutil.ReadFile("configs/recipe_graph.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &recipeGraphJson)
	if err != nil {
		log.Fatal(err)
	}

	// 第一遍遍历，创建一级Item，分配地址
	for key, jsonItem := range recipeGraphJson {
		(*Recipe)[key] = &RecipeGraphItem{
			ID:      int(jsonItem["id"].(float64)),
			Product: make([]recipeGraphProductItem, len(jsonItem["product"].([]interface{}))),
			Cost:    make([]interface{}, len(jsonItem["cost"].([]interface{}))),
		}
	}

	// 第二遍遍历
	for key, jsonItem := range recipeGraphJson {
		jsonProduct := jsonItem["product"].([]interface{})
		for idx, jsonProductItem := range jsonProduct {
			recipeItem := recipeGraphProductItem{}
			recipeItem.Count = int(jsonProductItem.(map[string]interface{})["count"].(float64))
			recipeItem.Time = jsonProductItem.(map[string]interface{})["time"].(float64)
			delete(jsonProductItem.(map[string]interface{}), "count")
			delete(jsonProductItem.(map[string]interface{}), "time")
			recipeItem.Items = make([]recipeGraphItemWithCount, len(jsonProductItem.(map[string]interface{})))
			itemIdx := 0

			// 赋值 ItemWithCount
			for jsonProductItemKey, jsonProductItemCount := range jsonProductItem.(map[string]interface{}) {
				itemWithCount := recipeGraphItemWithCount{
					Count: int(jsonProductItemCount.(float64)),
					Name:  jsonProductItemKey,
					Item:  (*Recipe)[jsonProductItemKey],
				}
				recipeItem.Items[itemIdx] = itemWithCount
				itemIdx++
			}

			(*Recipe)[key].Product[idx] = recipeItem
		}
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

	recipeGraphJson := make(map[string]interface{})
	var id int
	// 1001 - 1999 组件
	for id, i, j := 1001, 0, len(components); i < j; i, id = i+1, id+1 {
		name := components[i].Name
		gen := map[string]interface{}{
			"id":      id,
			"product": make([]map[string]interface{}, 0),
			"cost":    make([]map[string]interface{}, 0),
		}
		recipeGraphJson[name] = &gen
	}
	if id > 1999 {
		log.Fatal("component id greater than 1999 !")
	}

	// 2001 - 2999 建筑
	for id, i, j := 2001, 0, len(architectures); i < j; i, id = i+1, id+1 {
		name := architectures[i].Name
		gen := map[string]interface{}{
			"id":      id,
			"product": make([]map[string]interface{}, 0),
			"cost":    make([]map[string]interface{}, 0),
		}
		recipeGraphJson[name] = &gen
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

			recipeItemMap := recipeGraphJson[k].(*map[string]interface{})
			innerProduct := append((*recipeItemMap)["product"].([]map[string]interface{}), item)
			(*recipeItemMap)["product"] = &innerProduct
			recipeGraphJson[k] = recipeItemMap
		}
	}

	// 写入文件
	res, err := json.MarshalIndent(&recipeGraphJson, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("configs/recipe_graph.json", res, 0644); err != nil {
		log.Fatal(err)
	}
}

// ProductCalculate 产量计算
func (r *RecipeGraphItem) ProductCalculate(count int) []map[string]float64 {
	product := r.Product
	resultMapList := make([]map[string]float64, len(product))

	for formulaIndex, formulaItem := range product {
		resultMapList[formulaIndex] = make(map[string]float64)
		for _, item := range formulaItem.Items {
			// 加上当前元素产能
			resultMapList[formulaIndex][item.Name] += float64(count) * (float64(item.Count) / formulaItem.Time)
			// 合并当前元素子元素产能
			childMapList := item.Item.ProductCalculate(count * item.Count)
			if len(childMapList) > 0 {
				util.MergeMap(resultMapList[formulaIndex], childMapList[0])
			}
		}
	}

	return resultMapList
}

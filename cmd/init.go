/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/smokecat/dyson-sphere-program-tool/internal/gamedata"
	"github.com/spf13/cobra"
)

type genItem struct {
	Id int `json:"id"`
	Product []map[string]interface{} `json:"product"`
	Cost []map[string]interface{} `json:"cost"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: initCmdRun,
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initCmdRun(cmd *cobra.Command, args []string) {
	data, err := gamedata.LoadGameData()
	if err != nil {
		log.Fatal(err)
	}

	components := data.Component
	architectures := data.Architecture
	blueprints := data.Blueprint

	genJson := make(map[string]*genItem)
	var id int
	for id, i, j := 1001, 0, len(components); i < j; i, id = i + 1, id + 1 {
		name := components[i].Name
		gen := genItem{
			Id: id,
			Product: make([]map[string]interface{}, 0),
			Cost: make([]map[string]interface{}, 0),
		}
		genJson[name] = &gen
	}
	if id > 1999 {
		log.Fatal("component id greater than 1999 !")
	}

	for id, i, j := 2001, 0, len(architectures); i < j; i, id = i + 1, id + 1 {
		name := architectures[i].Name
		gen := genItem{
			Id: id,
			Product: make([]map[string]interface{}, 0),
			Cost: make([]map[string]interface{}, 0),
		}
		genJson[name] = &gen
	}
	if id > 2999 {
		log.Fatal("architecture id greater than 2999!")
	}

	for id, i, j := 3001, 0, len(blueprints); i < j; i, id = i + 1, id + 1 {
		time := blueprints[i].Time
		p := blueprints[i].Product
		c := blueprints[i].Cost

		// 遍历 product
		for k, v := range p {

			item := make(map[string]interface{})
			item["time"] = time
			item["count"] = v

			// 遍历 cost
			for ik, iv := range c {
				item[ik] = iv
			}

			genJson[k].Product = append(genJson[k].Product, item)
		}
	}




	// 写入文件
	res, err := json.MarshalIndent(&genJson, "", "  ")
	ioutil.WriteFile("configs/gen.json", res, 0644)
}

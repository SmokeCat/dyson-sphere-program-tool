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
	"fmt"
	"log"
	"strconv"

	"github.com/smokecat/dyson-sphere-program-tool/internal/gamedata"
	"github.com/spf13/cobra"
)

// productCmd represents the product command
var productCmd = &cobra.Command{
	Use:   "product",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: productCmdRun,
}

func init() {
	rootCmd.AddCommand(productCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// productCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// productCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func productCmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("product called ")

	// 获取参数，bluepoint idz
	if len(args) > 1 {
		log.Fatal("product: expect one args but have greater than one")
	}
	bpId, err :=  strconv.Atoi(args[0])
	if err != nil {
		log.Fatal((err))
	}

	// 获取配置
	data, err := gamedata.LoadGameData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("game version: ", data.Version)

	// 遍历消耗
	data.PrintCost(bpId)
}

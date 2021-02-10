package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/smokecat/dyson-sphere-program-tool/internal/gamedata"
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
}

func productCmdRun(cmd *cobra.Command, args []string) {
	// 获取参数，bluepoint idz
	if len(args) < 1 {
		log.Fatal("缺少参数: 需要指定物品")
	}
	target := args[0]
	count := 1
	if len(args) > 1{
		count, _ = strconv.Atoi(args[1])
	}

	fmt.Println("生产查询: ", target)

	// 获取配方
	recipeGraph, err := gamedata.LoadRecipeGraph()
	if err != nil {
		log.Fatal(err)
	}
	targetItem := (*recipeGraph)[target]
	if targetItem == nil {
		log.Fatal("找不到物品: ", target)
	}

	fmt.Println(targetItem.ProductCalculate(count))

}

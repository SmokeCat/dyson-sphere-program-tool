package cmd

import (
	"fmt"
	"log"

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
}

func productCmdRun(cmd *cobra.Command, args []string) {
	// 获取参数，bluepoint idz
	if len(args) != 1 {
		log.Fatal("缺少参数: 需要指定一个物品")
	}
	target := args[0]

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

	// 递归遍历
	res := targetItem.RecurCost()
	for i, j := 0, len(res); i < j; i++ {
		fmt.Printf("方案%v(产能:%g/s): \n", i+1, (*targetItem).Product[0]["count"].(float64)/(*targetItem).Product[0]["time"].(float64))
		for k, v := range res[i] {
			fmt.Printf("%v : %.2f/s\n", k, v)
		}
	}

}

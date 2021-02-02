package cmd

import (
	"github.com/smokecat/dyson-sphere-program-tool/internal/gamedata"
	"github.com/spf13/cobra"
)

type genItem struct {
	ID      int                      `json:"id"`
	Product []map[string]interface{} `json:"product"`
	Cost    []map[string]interface{} `json:"cost"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化游戏数据",
	Long:  `初始化游戏数据：长介绍`,
	Run:   initCmdRun,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initCmdRun(cmd *cobra.Command, args []string) {
	gamedata.InitRecipeGraph()
}

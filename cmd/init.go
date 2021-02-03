package cmd

import (
	"github.com/spf13/cobra"

	"github.com/smokecat/dyson-sphere-program-tool/internal/gamedata"
)

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

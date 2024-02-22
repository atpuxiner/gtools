package cmd

import (
	"github.com/atpuxiner/gtools/gtcli/internal/gtcmd/grapi"
	"github.com/spf13/cobra"
)

var (
	// grapiCmd represents the grapi command
	grapiCmd = &cobra.Command{
		Use:   "grapi",
		Short: "gin restful api",
		Long:  `grapi: gin-restful-api，集成gin、viper、zap、gorm...`,
	}
	// grapiNewCmd
	grapiNewCmd = &cobra.Command{
		Use:   "new",
		Short: "new grapi",
		Long:  "new grapi",
		Run: func(cmd *cobra.Command, args []string) {
			grapi.NewRun(cmd, args)
		},
	}
	// grapiAddCmd
	grapiAddCmd = &cobra.Command{
		Use:   "add",
		Short: "add api",
		Long:  "add api",
		Run: func(cmd *cobra.Command, args []string) {
			grapi.AddRun(cmd, args)
		},
	}
)

func init() {
	grapiCmd.AddCommand(grapiNewCmd)
	grapiCmd.AddCommand(grapiAddCmd)
	rootCmd.AddCommand(grapiCmd)
	// add flags
	// :: grapiNewCmd
	grapiNewCmd.Flags().StringP("proj", "p", "", "项目名称")
	grapiNewCmd.Flags().StringP("mod", "m", "", "模块名称")
	grapiNewCmd.Flags().StringP("dir", "d", ".", "目录")
	_ = grapiNewCmd.MarkFlagRequired("proj")
	_ = grapiNewCmd.MarkFlagRequired("mod")
	// :: grapiAddCmd
	grapiAddCmd.Flags().StringP("api", "a", "", "api名称")
	grapiAddCmd.Flags().StringP("ver", "v", "v1", "版本号")
	_ = grapiAddCmd.MarkFlagRequired("api")
}

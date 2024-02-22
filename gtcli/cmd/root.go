package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gtcli",
	Short: "gtcli: a command line for gtools",
	Long:  `gtcli: a command line for gtools`,
	Version: func() string {
		allVersion := []string{
			`1.0.0@初始版本: grapi`,
		}
		return strings.Split(allVersion[len(allVersion)-1], "@")[0]
	}(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// add flags
	rootCmd.Flags().BoolP("version", "v", false, "version")
}

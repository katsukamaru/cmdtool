package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

var RootCmd = &cobra.Command{
	Use:   "cmdtool",
	Short: "cmdtool",
	Long:  "Command tool",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cmdtool",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmdtool v" + VERSION)
	},
}

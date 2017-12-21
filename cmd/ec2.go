package cmd

import (
	"github.com/katsukamaru/cmdtool/cmd/ec2"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ec2Cmd)
}

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "List of instances.",
	Long:  "List of instances.",
	Run: func(cmd *cobra.Command, args []string) {
		ec2.List()
	},
}

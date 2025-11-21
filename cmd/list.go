package cmd

import (
	"github.com/skarekroe666/secepp/internal"
	"github.com/spf13/cobra"
)

var listAllCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the stored secrets",
	Run: func(cmd *cobra.Command, args []string) {
		internal.ListAllFiles()
	},
}

func init() {
	rootCmd.AddCommand(listAllCmd)
}

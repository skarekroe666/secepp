package cmd

import (
	"github.com/skarekroe666/secepp/internal"
	"github.com/spf13/cobra"
)

var getSecretByName = &cobra.Command{
	Use:   "get",
	Short: "get a specific secret",
	Run: func(cmd *cobra.Command, args []string) {
		internal.GetFile()
	},
}

func init() {
	rootCmd.AddCommand(getSecretByName)
}

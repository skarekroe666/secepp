package cmd

import (
	"github.com/skarekroe666/secepp/internal"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a secret",
	Run: func(cmd *cobra.Command, args []string) {
		internal.DeleteSecret()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

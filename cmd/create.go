package cmd

import (
	"github.com/skarekroe666/secepp/internal"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "create",
	Short: "create a secret for a file",
	Run: func(cmd *cobra.Command, args []string) {
		// fileName := ""
		// if len(args) == 1 {
		// 	fileName = args[0]
		// }
		internal.CreateSecret()
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
}

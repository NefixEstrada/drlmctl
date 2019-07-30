package cmd

import (
	"github.com/brainupdaters/drlmctl/cli"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout",
	Long:  `logout`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

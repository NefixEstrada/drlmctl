package cmd

import (
	"github.com/brainupdaters/drlmctl/cli"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "TODO",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Login(cmd.Flag("username").Value.String(), cmd.Flag("password").Value.String())
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Username of the new user")
	loginCmd.Flags().StringP("password", "p", "", "Password of the new user")

	rootCmd.AddCommand(loginCmd)
}

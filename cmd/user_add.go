package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/user"
	"github.com/brainupdaters/drlmctl/models"

	"github.com/spf13/cobra"
)

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user to DRLM",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		usr := &models.User{
			Username: cmd.Flag("username").Value.String(),
			Password: cmd.Flag("password").Value.String(),
		}

		user.Add(usr)
	},
}

func init() {
	userAddCmd.Flags().StringP("username", "u", "", "Username of the new user")
	userAddCmd.MarkFlagRequired("username")
	userAddCmd.Flags().StringP("password", "p", "", "Password of the new user")
	userAddCmd.MarkFlagRequired("password")

	userCmd.AddCommand(userAddCmd)
}

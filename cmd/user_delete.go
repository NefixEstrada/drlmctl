// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/user"

	"github.com/spf13/cobra"
)

var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an user from DRLM",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		user.Delete(cmd.Flag("username").Value.String())
	},
}

func init() {
	userDeleteCmd.Flags().StringP("username", "u", "", "Username of the user to be deleted")
	userDeleteCmd.MarkFlagRequired("username")

	userCmd.AddCommand(userDeleteCmd)
}

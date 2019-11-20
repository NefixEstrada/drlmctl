// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"os"

	"github.com/brainupdaters/drlmctl/cli/user"

	"github.com/spf13/cobra"
)

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the DRLM users",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		user.List(os.Stdout)
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
}

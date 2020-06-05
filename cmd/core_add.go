// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/core"

	"github.com/spf13/cobra"
)

var coreAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new DRLM Core host and copy the SSH keys",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: STDIN
		core.Add(fs, coreHostFlag)
	},
}

var coreHostFlag string

func init() {
	coreAddCmd.Flags().StringVarP(&coreHostFlag, "host", "", "", "Hostname / IP of the server where DRLM Core is going to be running")
	coreAddCmd.MarkFlagRequired("host")

	coreCmd.AddCommand(coreAddCmd)
}

// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/plugin"

	"github.com/spf13/cobra"
)

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available DRLM Plugins",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		plugin.List(pluginListRepoFlag)
	},
}

var pluginListRepoFlag string

func init() {
	pluginListCmd.Flags().StringVarP(&pluginListRepoFlag, "repo", "r", "", "specific repository to list plugins")

	pluginRepoCmd.AddCommand(pluginListCmd)
}

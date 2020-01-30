// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/plugin"

	"github.com/spf13/cobra"
)

var pluginAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a plugin to a DRLM Agent",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		plugin.Add(pluginAddHostFlag, pluginAddPluginFlag, pluginAddVersionFlag)
	},
}

var (
	pluginAddHostFlag,
	pluginAddPluginFlag,
	pluginAddVersionFlag string
)

func init() {
	pluginAddCmd.Flags().StringVarP(&pluginAddHostFlag, "host", "", "", "host of the agent where the plugin is going to be installed")
	pluginAddCmd.MarkFlagRequired("host")
	pluginAddCmd.Flags().StringVarP(&pluginAddPluginFlag, "plugin", "p", "", "plugin that is going to be installed")
	pluginAddCmd.MarkFlagRequired("plugin")
	pluginAddCmd.Flags().StringVarP(&pluginAddVersionFlag, "version", "", "", "version of the plugin that is going to be installed")

	pluginRepoCmd.AddCommand(pluginAddCmd)
}

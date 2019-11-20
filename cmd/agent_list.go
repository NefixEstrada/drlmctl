// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/agent"

	"github.com/spf13/cobra"
)

var agentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the DRLM Agents",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		agent.List()
	},
}

func init() {
	agentCmd.AddCommand(agentListCmd)
}

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/agent/request"

	"github.com/spf13/cobra"
)

var agentRequestListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the DRLM Agents that have requested to join the DRLM Core",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		request.List()
	},
}

func init() {
	agentRequestCmd.AddCommand(agentRequestListCmd)
}

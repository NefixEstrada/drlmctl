package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/agent/request"

	"github.com/spf13/cobra"
)

var agentRequestAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "Accept an Agent to DRLM",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		request.Accept(agentRequestAcceptHost)
	},
}

var (
	agentRequestAcceptHost string
)

func init() {
	agentRequestAcceptCmd.Flags().StringVarP(&agentRequestAcceptHost, "host", "", "", "Host of the agent to accept")
	agentRequestAcceptCmd.MarkFlagRequired("host")

	agentRequestCmd.AddCommand(agentRequestAcceptCmd)
}

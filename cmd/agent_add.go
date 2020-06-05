// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/agent"

	"github.com/spf13/cobra"
)

var agentAddCmd = &cobra.Command{
	Use:   "add",
	Short: `Add a new DRLM Agent and copy the SSH keys`,
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: STDIN
		agent.Add(agentHostFlag)
	},
}

var agentHostFlag string
var agentPortFlag int
var agentUsrFlag string
var agentPwdFlag string
var agentIsAdminFlag bool

func init() {
	agentAddCmd.Flags().StringVarP(&agentHostFlag, "host", "", "", "Hostname / IP of the server where the DRLM Agent is going to be running")
	agentAddCmd.MarkFlagRequired("host")

	agentCmd.AddCommand(agentAddCmd)
}

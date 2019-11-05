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
		agent.Add(agentHostFlag, agentPortFlag, agentUsrFlag, agentPwdFlag)
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
	agentAddCmd.Flags().IntVarP(&agentPortFlag, "port", "", 22, "SSH Port of the host")
	agentAddCmd.Flags().StringVarP(&agentUsrFlag, "user", "u", "", "SSH Username")
	agentAddCmd.MarkFlagRequired("user")
	agentAddCmd.Flags().StringVarP(&agentPwdFlag, "password", "p", "", "SSH Password. If the parameter isn't provided, it's going to be asked through stdin")

	agentCmd.AddCommand(agentAddCmd)
}

// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/agent"

	"github.com/spf13/cobra"
)

var agentInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Compile / Download the DRLM Agent binary, install and start it on the server",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		agent.Install(fs, agentHostFlag, agentPortFlag, agentUsrFlag, agentPwdFlag, agentVersionFlag)
	},
}

var agentVersionFlag string

func init() {
	agentInstallCmd.Flags().StringVarP(&agentHostFlag, "host", "", "", "Host of the DRLM Agent")
	agentInstallCmd.MarkFlagRequired("host")
	agentAddCmd.Flags().IntVarP(&agentPortFlag, "port", "", 22, "SSH Port of the host")
	agentAddCmd.Flags().StringVarP(&agentUsrFlag, "user", "u", "", "SSH Username")
	agentAddCmd.MarkFlagRequired("user")
	agentAddCmd.Flags().StringVarP(&agentPwdFlag, "password", "p", "", "SSH Password. If the parameter isn't provided, it's going to be asked through stdin")

	agentInstallCmd.Flags().StringVarP(&agentVersionFlag, "version", "", "", "Git tag of the version of DRLM Agent")

	agentCmd.AddCommand(agentInstallCmd)
}

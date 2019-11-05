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
		agent.Install(agentHostFlag, agentVersionFlag)
	},
}

var agentVersionFlag string

func init() {
	agentInstallCmd.Flags().StringVarP(&agentHostFlag, "host", "", "", "Host of the DRLM Agent")
	agentInstallCmd.MarkFlagRequired("id")

	agentInstallCmd.Flags().StringVarP(&agentVersionFlag, "version", "", "", "Git tag of the version of DRLM Agent")

	agentCmd.AddCommand(agentInstallCmd)
}

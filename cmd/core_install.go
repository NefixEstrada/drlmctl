package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/core"

	"github.com/spf13/cobra"
)

var coreInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Compile / Download the DRLM Core binary, install and start it on the server",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Install(versionFlag)
	},
}

var versionFlag string

func init() {
	coreInstallCmd.Flags().StringVarP(&versionFlag, "version", "", "", "Git tag of the version of DRLM Core")

	coreCmd.AddCommand(coreInstallCmd)
}

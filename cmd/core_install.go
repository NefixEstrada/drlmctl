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
		core.Install(coreVersionFlag)
	},
}

var coreVersionFlag string

func init() {
	coreInstallCmd.Flags().StringVarP(&coreVersionFlag, "version", "", "", "Git tag of the version of DRLM Core")

	coreCmd.AddCommand(coreInstallCmd)
}

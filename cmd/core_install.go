// SPDX-License-Identifier: AGPL-3.0-only

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
		core.Install(fs, coreVersionFlag, corePortFlag, coreUsrFlag, corePwdFlag)
	},
}

var (
	coreVersionFlag string
	corePortFlag    int
	coreUsrFlag     string
	corePwdFlag     string
	coreIsAdminFlag bool
)

func init() {
	coreInstallCmd.Flags().StringVarP(&coreVersionFlag, "version", "", "", "Git tag of the version of DRLM Core")
	coreInstallCmd.Flags().IntVarP(&corePortFlag, "port", "", 22, "SSH Port of the host")
	coreInstallCmd.Flags().StringVarP(&coreUsrFlag, "user", "u", "", "SSH Username")
	coreInstallCmd.MarkFlagRequired("host")
	coreInstallCmd.Flags().StringVarP(&corePwdFlag, "password", "p", "", "SSH Password. If the parameter isn't provided, it's going to be asked through stdin")

	coreCmd.AddCommand(coreInstallCmd)
}

// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/core"

	"github.com/spf13/cobra"
)

var coreAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new DRLM Core host and copy the SSH keys",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: STDIN
		core.Add(coreHostFlag, corePortFlag, coreUsrFlag, corePwdFlag)
	},
}

var coreHostFlag string
var corePortFlag int
var coreUsrFlag string
var corePwdFlag string
var coreIsAdminFlag bool

func init() {
	coreAddCmd.Flags().StringVarP(&coreHostFlag, "host", "", "", "Hostname / IP of the server where DRLM Core is going to be running")
	coreAddCmd.MarkFlagRequired("host")
	coreAddCmd.Flags().IntVarP(&corePortFlag, "port", "", 22, "SSH Port of the host")
	coreAddCmd.Flags().StringVarP(&coreUsrFlag, "user", "u", "", "SSH Username")
	coreAddCmd.MarkFlagRequired("user")
	coreAddCmd.Flags().StringVarP(&corePwdFlag, "password", "p", "", "SSH Password. If the parameter isn't provided, it's going to be asked through stdin")

	coreCmd.AddCommand(coreAddCmd)
}

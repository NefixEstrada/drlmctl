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
		core.Add(hostFlag, portFlag, usrFlag, pwdFlag, isAdminFlag)
	},
}

var hostFlag string
var portFlag int
var usrFlag string
var pwdFlag string
var isAdminFlag bool

func init() {
	coreAddCmd.Flags().StringVarP(&hostFlag, "host", "", "", "Hostname / IP of the server where DRLM Core is going to be running")
	coreAddCmd.MarkFlagRequired("host")
	coreAddCmd.Flags().IntVarP(&portFlag, "port", "", 22, "SSH Port of the host")
	coreAddCmd.Flags().StringVarP(&usrFlag, "user", "u", "", "SSH Username")
	coreAddCmd.MarkFlagRequired("user")
	coreAddCmd.Flags().StringVarP(&pwdFlag, "password", "p", "", "SSH Password. If the parameter isn't provided, it's going to be asked through stdin")
	coreAddCmd.Flags().BoolVarP(&isAdminFlag, "is-admin", "a", false, "Sets if the user is admin (and thus doesn't require to upgrade their privileges to superuser) or if it's not (and thus it does require a privileges upgrade")

	coreCmd.AddCommand(coreAddCmd)
}

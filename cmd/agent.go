// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import "github.com/spf13/cobra"

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "TODO",
	Long:  `TODO`,
}

func init() {
	rootCmd.AddCommand(agentCmd)
}

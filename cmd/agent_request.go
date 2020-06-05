package cmd

import "github.com/spf13/cobra"

var agentRequestCmd = &cobra.Command{
	Use:   "request",
	Short: "TODO",
	Long:  `TODO`,
}

func init() {
	agentCmd.AddCommand(agentRequestCmd)
}

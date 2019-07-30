package cmd

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "TODO",
	Long:  `TODO`,
}

func init() {
	rootCmd.AddCommand(userCmd)
}

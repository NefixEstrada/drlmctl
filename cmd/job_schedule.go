// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cli/job"

	"github.com/spf13/cobra"
)

var jobScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedules a new job",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		job.Schedule(jobHostFlag, jobNameFlag, jobConfigFlag)
	},
}

var jobHostFlag string
var jobNameFlag string
var jobConfigFlag string

func init() {
	jobScheduleCmd.Flags().StringVarP(&jobHostFlag, "host", "", "", "Host of the DRLM Agent")
	jobScheduleCmd.MarkFlagRequired("host")
	jobScheduleCmd.Flags().StringVarP(&jobNameFlag, "name", "n", "", "Name of the job that needs to be scheduled")
	jobScheduleCmd.MarkFlagRequired("name")
	jobScheduleCmd.Flags().StringVarP(&jobConfigFlag, "config", "c", "", "The configuration that will be passed to the plugin")
	jobScheduleCmd.MarkFlagRequired("config")

	jobCmd.AddCommand(jobScheduleCmd)
}

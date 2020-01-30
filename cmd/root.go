// SPDX-License-Identifier: AGPL-3.0-only

package cmd

import (
	"github.com/brainupdaters/drlmctl/cfg"
	"github.com/brainupdaters/drlmctl/core"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	logger "github.com/brainupdaters/drlm-common/pkg/log"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string
var logVerbose bool

var rootCmd = &cobra.Command{
	Use:   "drlmctl",
	Short: "TODO",
	Long:  `TODO`,
}

// Execute is the main function of the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "configuration file to use instead of the default ($HOME/.drlmctl.toml)")
	rootCmd.PersistentFlags().BoolVarP(&logVerbose, "verbose", "v", false, "verbose logging output")
}

func initConfig() {
	fs.Init()
	cfg.Init(cfgFile)
	logger.Init(cfg.Config.Log)
	core.Init()
}

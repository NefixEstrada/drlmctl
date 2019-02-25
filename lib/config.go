package lib

import (
	"fmt"
	"os"
	"strings"

	"github.com/brainupdaters/drlm-common/logger"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type DrlmcliConfig struct {
	Drlmcore DrlmcoreConfig       `mapstructure:"drlmcore"`
	Logging  logger.LoggingConfig `mapstructure:"logging"`
}

var Config *DrlmcliConfig

func InitConfig(c string) {
	if c != "" {
		// Use config file from the flag.
		viper.SetConfigFile(c)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".drlm-core" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".drlm-cli")
	}

	// Enable environment variables
	// ex.: DRLMCLI_DRLMCORE_PORT=8000
	viper.SetEnvPrefix("DRLMCLI")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// var config Config

	err := viper.Unmarshal(&Config)
	if err != nil {
		panic("Unable to unmarshal config")
	}
}

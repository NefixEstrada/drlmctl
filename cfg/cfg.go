package cfg

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/brainupdaters/drlm-common/pkg/fs"
	logger "github.com/brainupdaters/drlm-common/pkg/log"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config has the values of the user configuration
var Config *DRLMCtlConfig

// DRLMCtlConfig is the configuration of the CLI tool of DRLM
type DRLMCtlConfig struct {
	Core DRLMCtlCoreConfig `mapstructure:"core"`
	Log  logger.Config     `mapstructure:"log"`
}

// DRLMCtlCoreConfig is the configuration related with the DRLM Core of the CLI tool of DRLM
type DRLMCtlCoreConfig struct {
	Host          string    `mapstructure:"host"`
	Port          int       `mapstructure:"port"`
	TLS           bool      `mapstructure:"tls"`
	CertPath      string    `mapstructure:"cert_path"`
	Tkn           string    `mapstructure:"tkn"`
	TknExpiration time.Time `mapstructure:"tkn_expiration"`
}

// v is the viper instance for the configuration
var v *viper.Viper

// Init prepares the configuration and reads it
func Init(cfgFile string) {
	v = viper.New()
	v.SetFs(fs.FS)
	SetDefaults()

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	}

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error reading the configuration: %v", err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		log.Fatalf("error decoding the configuration: invald configuration: %v", err)
	}
}

// SetDefaults sets the default configurations for Viper
func SetDefaults() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("error getting the home directory: %v", err)
	}

	v.SetConfigName("drlmctl")
	v.AddConfigPath(".")
	v.AddConfigPath(filepath.Join(home, ".drlm"))
	v.AddConfigPath(filepath.Join(home, ".config/drlm"))
	v.AddConfigPath("/etc/drlm")

	v.SetDefault("core", map[string]interface{}{
		"host":           "localhost",
		"port":           50051,
		"tls":            true,
		"cert_path":      "cert/server.crt",
		"tkn":            "",
		"tkn_expiration": nil,
	})

	v.SetDefault("log", map[string]interface{}{
		"level": "info",
		"file":  filepath.Join(home, ".log/drlm/drlmctl.log"),
	})

	v.SetEnvPrefix("DRLMCTL")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
}

// SaveTkn saves the token to the configuration file
func SaveTkn(tkn string, expiresAt time.Time) error {
	Config.Core.Tkn = tkn
	Config.Core.TknExpiration = expiresAt

	v.Set("core.tkn", tkn)
	v.Set("core.tkn_expiration", expiresAt)

	return v.WriteConfig()
}

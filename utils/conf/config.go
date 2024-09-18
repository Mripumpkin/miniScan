package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// Settings holds the application configuration
var Settings = &SettingsConf{}

// SettingsConf contains configuration settings
type SettingsConf struct {
	App App `mapstructure:"app"`
}

// App holds application-specific settings
type App struct {
	Secret string `mapstructure:"jwt_secret_token"`
	Salt   string `mapstructure:"salt"`
}

// Provider defines a set of read-only methods for accessing the application
// configuration parameters as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var (
	defaultConfig *viper.Viper
	provider      *viper.Viper
	once          sync.Once
)

// Config returns the default config provider
func Config() Provider {
	defaultConfig = readViperConfig()
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider() Provider {
	once.Do(func() {
		provider = readViperConfig()
	})
	return provider
}

func init() {
	defaultConfig = readViperConfig()
}

// readViperConfig reads and returns a configured viper instance
func readViperConfig() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix("miniScan")
	v.AddConfigPath(".")
	v.AddConfigPath("./")
	v.SetConfigName("config")
	v.SetConfigType("toml")

	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ratal error config file: %s ", err))
	}

	// Set default values
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	if err := v.Unmarshal(Settings); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %v", err))
	}

	// Load local development configuration if applicable
	runLevel := v.GetString("run_level")
	if runLevel == "development" {
		v.SetConfigName("config.local")
		v.SetConfigType("toml")
		v.AddConfigPath(".")
		v.MergeInConfig()
	}

	return v
}

package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var Settings = &SettingsConf{}

type SettingsConf struct {
	App App `mapstructure:"app"`
}

type App struct {
	Secret string `mapstructure:"jwt_secret_token"`
	Salt   string `mapstructure:"salt"`
}

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
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

var defaultConfig *viper.Viper
var provider *viper.Viper
var once sync.Once
var onceRedis sync.Once

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider() Provider {

	once.Do(func() {
		provider = readViperConfig()
	})
	return provider
}

func LoadConfigProviderRedis() Provider {
	onceRedis.Do(func() {
		provider = readViperConfig()
	})
	return provider
}

func init() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	// 环境变量设置支持
	v := viper.New()
	v.SetEnvPrefix("op")

	// 文件设置支持
	v.AddConfigPath(".")
	v.AddConfigPath("/srv/op_dev")
	v.SetConfigName("settings")
	v.SetConfigType("toml")

	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ratal error config file: %s ", err))
	}

	// global defaults
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")

	_ = v.Unmarshal(Settings)

	// 本地开发配置
	runLevel := v.GetString("run_level")
	if runLevel == "development" {
		v.SetConfigName("settings.local")
		v.SetConfigType("toml")
		v.AddConfigPath(".")
		v.MergeInConfig()
	}

	return v
}

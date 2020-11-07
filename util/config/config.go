package config

import (
	"github.com/spf13/viper"
	"strings"
)

type FileType int

const (
	Yaml FileType = iota
	Json
)

func (t FileType) String() string {
	switch t {
	case Yaml:
		return "yaml"
	case Json:
		return "json"
	default:
		return "yaml"
	}
}

func Get(_type FileType, name string, paths ...string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(_type.String())
	config.SetConfigName(name)
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()
	for _, path := range paths {
		config.AddConfigPath(path)
	}
	return config
}

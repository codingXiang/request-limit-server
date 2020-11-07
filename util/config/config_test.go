package config

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	configType = "yaml"
	configName = "test"
	configPath = []string{".", "./config"}
)

//Suite struct
type Suite struct {
	suite.Suite
	config *viper.Viper
}

//初始化 Suite
func (s *Suite) SetupSuite() {
	config := viper.New()
	config.SetConfigName(configType)
	config.SetConfigName(configName)
	for _, path := range configPath {
		config.AddConfigPath(path)
	}
	s.config = config
}

//TestStart 為測試程式進入點
func TestStartSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestGet() {
	test := Get(Yaml, configName, configPath...)
	assert.Equal(s.T(), test, s.config)
}

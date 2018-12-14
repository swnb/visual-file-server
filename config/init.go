package config

import (
	"strings"

	"github.com/spf13/viper"
)

const configDir = "cfm/"

var envMap map[string]string

var (
	baseConfig, selfConfig *viper.Viper
)

func init() {
	baseConfig = baseConfInit()

	envMap = resolveEnv() // 当前的.env的配置信息解析,最优先的

	selfConfig = selfConfInit(envMap)
}

func baseConfInit() *viper.Viper {
	baseConfig = viper.New()
	baseConfig.SetConfigFile(configDir + "base/config.yml")
	err := baseConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return baseConfig
}

func selfConfInit(envMap map[string]string) *viper.Viper {
	selfConfig := viper.New()
	if runMode, ok := envMap["runMode"]; ok {
		selfConfig.SetConfigFile(configDir + strings.ToLower(runMode) + "/config.yml")
	} else {
		selfConfig.SetConfigFile(configDir + "/prod/config.yml")
	}

	err := selfConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	for k, v := range envMap {
		selfConfig.Set(k, v)
	}

	return selfConfig
}

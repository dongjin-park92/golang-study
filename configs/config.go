package configs

import "github.com/spf13/viper"

type ViperConfig struct {
	*viper.Viper
}

var configs *ViperConfig

func init() {
	configs = readConfig()
}

func GetConfig() *ViperConfig {
	return configs
}

// Viper 패키지를 통해 config.yaml 파일 기반으로 환경설정값 읽기
func readConfig() *ViperConfig {
	viperConfig := viper.New()
	viperConfig.AddConfigPath("./configs")
	viperConfig.SetConfigName("config")
	// Pod ENV를 읽을 수 있도록 설정
	viperConfig.AutomaticEnv()
	err := viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return &ViperConfig{
		Viper: viperConfig,
	}
}

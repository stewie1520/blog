package config

import "github.com/spf13/viper"

type Config struct {
	Port          int16  `mapstructure:"PORT"`
	GrpcPort      int16  `mapstructure:"GRPC_PORT"`
	ApiDomain     string `mapstructure:"API_DOMAIN"`
	WebsiteDomain string `mapstructure:"WEBSITE_DOMAIN"`
	GinMode       string `mapstructure:"GIN_MODE"`
	DatabaseURL   string `mapstructure:"DATABASE_URL"`
	Google        struct {
		ClientID     string `mapstructure:"CLIENT_ID"`
		ClientSecret string `mapstructure:"CLIENT_SECRET"`
	} `mapstructure:"GOOGLE"`
}

func Init() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

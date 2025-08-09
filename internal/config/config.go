package config

import "github.com/spf13/viper"

func Init(filepath string) error {
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

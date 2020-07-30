package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func InitViper() error {
	config := Configuration{}
	viper.SetEnvPrefix(applicationEnvPrefix)
	viper.SetConfigFile(configFile)
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Info("config not loaded correctly")
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.WithError(err).Info("config not loaded correctly")
		return err
	}
	return nil
}
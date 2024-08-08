package loadcfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type PostgresqlConfig struct {
	Host string `mapstructure:"hostname"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}
type RedisConfig struct {
	Host string `mapstructure:"hostname"`
}

type DatabaseConfig struct {
	Psql PostgresqlConfig `mapstructure:"postgresql"`
	Rds  RedisConfig      `mapstructure:"redis"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Db  DatabaseConfig `mapstructure:"database"`
	SRV ServerConfig   `mapstructure:"server"`
}

func LoadYamlCfg() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	var c Config

	if err := viper.ReadInConfig(); err != nil {
		return c, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		return c, fmt.Errorf("couldn't read config: %w", err)
	}

	return c, nil
}

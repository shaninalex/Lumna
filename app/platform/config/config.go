package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Interface interface {
	Env() Environment
	Int(param string) int
	String(param string) string
	Bool(param string) bool
	StringSlice(param string) []string
}

type Environment string

const (
	EnvironmentDev  Environment = "dev"
	EnvironmentTest Environment = "testing"
)

type Config struct {
	v *viper.Viper
}

func (s *Config) Env() Environment { return Environment(s.String("env")) }

func (s *Config) Int(param string) int { return s.v.GetInt(param) }

func (s *Config) String(param string) string { return s.v.GetString(param) }

func (s *Config) Bool(param string) bool { return s.v.GetBool(param) }

func (s *Config) StringSlice(param string) []string { return s.v.GetStringSlice(param) }

func ReadConfig(path string) *Config {
	s := &Config{
		v: viper.New(),
	}
	s.v.SetConfigFile(path)
	if err := s.v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Can't open config file. %s \n", err))
	}
	if err := s.v.Unmarshal(s); err != nil {
		panic(fmt.Errorf("Can't unmarshal config. %s \n", err))
	}
	return s
}

func ProvideConfig(configPath string) func() *Config {
	return func() *Config {
		return ReadConfig(configPath)
	}
}

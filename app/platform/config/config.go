package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gitlab.com/shaninalex/lumna/app/core/bus"
)

type Interface interface {
	Env() Environment
	Int(param string) int
	String(param string) string
	Bool(param string) bool
	StringSlice(param string) []string
	AuthSecret() string
	SetupEnabled() bool
	CORSOrigins() []string
	SecureCookies() bool
	EmbedSPA() bool
	EmailConfig() EmailerConfig
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

func (s *Config) AuthSecret() string { return s.v.GetString("secret_key") }

func (s *Config) SetupEnabled() bool { return s.v.GetBool("serve.setup") }

// CORSOrigins lists the origins allowed to call the API with credentials.
// Empty means same-origin only, which is the default self-hosted setup.
func (s *Config) CORSOrigins() []string { return s.v.GetStringSlice("serve.cors_origins") }

// SecureCookies must be false only for plain-http local development.
func (s *Config) SecureCookies() bool { return s.v.GetBool("serve.secure_cookies") }

// EmbedSPA include route with embedded SPA build or not
func (s *Config) EmbedSPA() bool { return s.v.GetBool("serve.embed_spa") }

// EmailConfig return email everything email sender need to send and process emails
func (s *Config) EmailConfig() EmailerConfig {
	return EmailerConfig{
		Process:  s.v.GetBool("smtp.process"),
		From:     s.v.GetString("smtp.from"),
		Host:     s.v.GetString("smtp.host"),
		Port:     s.v.GetInt("smtp.port"),
		User:     s.v.GetString("smtp.user"),
		Password: bus.Secret(s.v.GetString("smtp.password")),
	}
}

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

func ProvideConfig(configPath string) *Config {
	return ReadConfig(configPath)
}

type EmailerConfig struct {
	Process  bool
	From     string
	Host     string
	Port     int
	User     string
	Password bus.Secret
}

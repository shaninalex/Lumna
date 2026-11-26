package test

import (
	"github.com/spf13/cast"
	"gitlab.com/shaninalex/lumna/app/core/secure"
	"gitlab.com/shaninalex/lumna/app/platform/config"
)

var _ config.Interface = MockConfig{}

// MockConfig is a config.Interface
type MockConfig struct {
	data map[string]any
}

func NewMockConfig(data map[string]any) MockConfig {
	if data == nil {
		data = map[string]any{}
	}
	return MockConfig{data}
}

func (m MockConfig) Env() config.Environment { return config.Environment(m.String("env")) }

func (m MockConfig) Int(param string) int { return cast.ToInt(m.data[param]) }

func (m MockConfig) String(param string) string { return cast.ToString(m.data[param]) }

func (m MockConfig) Bool(param string) bool { return cast.ToBool(m.data[param]) }

func (m MockConfig) StringSlice(param string) []string { return cast.ToStringSlice(m.data[param]) }

func (m MockConfig) AuthSecret() string { return m.String("secret_key") }

func (m MockConfig) SetupEnabled() bool { return m.Bool("serve.setup") }

func (m MockConfig) SecureCookies() bool { return m.Bool("serve.secure_cookies") }

func (m MockConfig) EmbedSPA() bool { return m.Bool("serve.embed_spa") }

func (m MockConfig) EmailConfig() config.EmailerConfig {
	return config.EmailerConfig{
		From:      m.String("smtp.from"),
		Host:      m.String("smtp.host"),
		Port:      m.Int("smtp.port"),
		User:      m.String("smtp.user"),
		Password:  secure.Secret(m.String("smtp.password")),
		ServerUrl: m.String("smtp.server_url"),
	}
}

package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Service serviceConfig
}

type serviceConfig struct {
	AddressHTTP   string `envconfig:"ADDRESS_HTTP"`
	AddressHTTPS  string `envconfig:"ADDRESS_HTTPS"`
	PathCertHTTPS string `envconfig:"PATH_CERT_HTTPS"`
	PathKeyHTTPS  string `envconfig:"PATH_KEY_HTTPS"`

	ReadTimeout  int64 `envconfig:"READ_TIMEOUT"`
	WriteTimeout int64 `envconfig:"WRITE_TIMEOUT"`

	PathFileLogs string `envconfig:"PATH_FILE_LOGS"`
}

func ConfigFromEnv(prefix string) *Config {
	APP := serviceConfig{
		AddressHTTP:   "",
		AddressHTTPS:  "",
		PathCertHTTPS: "",
		ReadTimeout:   0,
		WriteTimeout:  0,
		PathFileLogs:  "",
	}

	_ = envconfig.Process(prefix, &APP)
	return &Config{
		Service: APP,
	}
}

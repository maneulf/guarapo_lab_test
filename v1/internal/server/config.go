package server

import "github.com/kelseyhightower/envconfig"

type config struct {
	AddressHTTP   string `envconfig:"ADDRESS_HTTP"`
	AddressHTTPS  string `envconfig:"ADDRESS_HTTPS"`
	PathCertHTTPS string `envconfig:"PATH_CERT_HTTPS"`
	PathKeyHTTPS  string `envconfig:"PATH_KEY_HTTPS"`

	ReadTimeout  int64 `envconfig:"READ_TIMEOUT"`
	WriteTimeout int64 `envconfig:"WRITE_TIMEOUT"`

	PathFileLogs string `envconfig:"PATH_FILE_LOGS"`

	PERSISTENCE_TYPE string `envconfig:"PERSISTENCE_TYPE"`
}

// ConfigFromEnv ...
func ConfigFromEnv() (*config, error) {
	conf := config{
		AddressHTTP:   ":80",
		AddressHTTPS:  ":443",
		PathCertHTTPS: "",
		PathKeyHTTPS:  "",

		ReadTimeout:  10,
		WriteTimeout: 10,

		PathFileLogs: "",
	}
	e := envconfig.Process("", &conf)
	return &conf, e
}

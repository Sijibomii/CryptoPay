package config

type ServerConfig struct {
	Host string
	Port int64
	// PrivateKeyPath string
	// PublicKeyPath  string
}

type Config struct {
	Postgres string
	Server   ServerConfig
}

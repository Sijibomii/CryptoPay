package config

type ServerConfig struct {
	Host string
	Port int64
	// PrivateKeyPath string
	// PublicKeyPath  string
}

type MailerConfig struct {
	SmtpHost     string
	SmtpPort     int
	SmtpUsername string
	SmtpPassword string
}

type Config struct {
	Postgres string
	Server   ServerConfig
	Mailer   MailerConfig
}

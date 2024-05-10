package config

type Config struct {
	Application Application `yaml:"application"`
	Adapters    Adapters    `yaml:"adapters"`
}

type Application struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Adapters struct {
	Primary   Primary   `yaml:"primary"`
	Secondary Secondary `yaml:"secondary"`
}

type Primary struct {
	HttpAdapter HttpAdapter `yaml:"httpAdapter"`
}

type Secondary struct {
	Databases Databases `yaml:"databases"`
}

// PROVIDERS CONFIG

type Databases struct {
	Messages Database `yaml:"messages"`
}

type Database struct {
	Type       string            `yaml:"type"`
	Host       string            `env:"MESSAGES_HOST"     env-required:"true" yaml:"host"`
	Port       string            `env:"MESSAGES_PORT"     env-required:"true" yaml:"port"`
	User       string            `env:"MESSAGES_USER"     env-required:"true" yaml:"user"`
	Password   string            `env:"MESSAGES_PASSWORD" env-required:"true" yaml:"password"`
	Name       string            `yaml:"name"`
	Procedures map[string]string `yaml:"procedures"`
}

// SERVER CONFIG

type HttpAdapter struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port string `yaml:"port"`
}

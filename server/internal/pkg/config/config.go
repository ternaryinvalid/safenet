package config

type Config struct {
	Databases Databases
}

type Databases struct {
	UserRepository Database
}

type Database struct {
	Type     string
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

func New() Config {
	return Config{}
}

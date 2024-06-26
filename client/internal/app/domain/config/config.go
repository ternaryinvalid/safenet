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
	Providers    Providers    `yaml:"providers"`
	Repositories Repositories `yaml:"repositories"`
}

type Repositories struct {
	Cache Cache `yaml:"cache"`
}

type Cache struct {
	Filepath string `yaml:"filepath"`
}

// PROVIDERS CONFIG

type Providers struct {
	ServerProvider ServerProvider `yaml:"serverProvider"`
}

type ServerProvider struct {
	Host      string                  `yaml:"host"`
	Endpoints ServerProviderEndpoints `yaml:"endpoints"`
}

type ServerProviderEndpoints struct {
	Get  Endpoint `yaml:"get"`
	Send Endpoint `yaml:"send"`
}

type Endpoint struct {
	Method  string            `yaml:"method"`
	Path    string            `yaml:"path"`
	Headers map[string]string `yaml:"headers"`
}

// SERVER CONFIG

type HttpAdapter struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port string `yaml:"port"`
}

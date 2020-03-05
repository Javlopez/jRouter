package jrouter

type Config struct {
	File string
}

func NewConfig() *Config {
	return &Config{}
}

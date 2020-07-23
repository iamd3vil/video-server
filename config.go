package main

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

// Config is the config given by the user
type Config struct {
	HTTPAddr string `koanf:"listen_addr"`
	Dest     string `koanf:"destination"`
}

func initConfig(path string) (Config, error) {
	cfg := Config{}
	k := koanf.New(".")

	if err := k.Load(file.Provider(path), toml.Parser()); err != nil {
		return Config{}, err
	}

	k.Unmarshal("config", &cfg)
	return cfg, nil
}

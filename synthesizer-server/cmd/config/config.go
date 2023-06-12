package config

import (
	"os"
	"strings"
)

const (
	base_openai = "base_openai"
	token       = "token"
	port        = "port"
)

type Config struct {
	BaseOpenAI string
	Port       string
	Token      string
}

func Make() Config {
	defaults := map[string]any{
		base_openai: "https://api.openai.com",
		token:       "",
		port:        ":8081",
	}

	for k := range defaults {
		if v, ok := os.LookupEnv(strings.ToUpper(k)); ok {
			defaults[k] = v
		}
	}
	return Config{
		BaseOpenAI: defaults[base_openai].(string),
		Token:      defaults[token].(string),
		Port:       defaults[port].(string),
	}
}

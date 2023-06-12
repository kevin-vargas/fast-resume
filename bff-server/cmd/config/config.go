package config

import (
	"os"
	"strings"
	"time"
)

const (
	base_synthesizer_server = "base_synthesizer_server"
	slack_api_uri           = "slack_api_uri"
	port                    = "port"
	timeout                 = "timeout"
	allow_origins           = "allow_origins"
	ttl_service             = "ttl_service"
	api_prefix              = "api_prefix"
)

type Config struct {
	BaseSynthesizerServer string
	SlackAPIUri           string
	Port                  string
	Timeout               time.Duration
	AllowOrigins          string
	APIPrefix             string
	TTLService            time.Duration
}

func Make() Config {
	defaults := map[string]any{
		base_synthesizer_server: "http://localhost:8081",
		slack_api_uri:           "https://slack.com/api",
		allow_origins:           "http://localhost,http://localhost:5173",
		ttl_service:             5 * 60,
		api_prefix:              "",
		timeout:                 8,
		port:                    ":8080",
	}

	for k := range defaults {
		if v, ok := os.LookupEnv(strings.ToUpper(k)); ok {
			defaults[k] = v
		}
	}
	return Config{
		BaseSynthesizerServer: defaults[base_synthesizer_server].(string),
		SlackAPIUri:           defaults[slack_api_uri].(string),
		Port:                  defaults[port].(string),
		AllowOrigins:          defaults[allow_origins].(string),
		APIPrefix:             defaults[api_prefix].(string),
		Timeout:               time.Duration(defaults[timeout].(int) * int(time.Second)),
		TTLService:            time.Duration(defaults[ttl_service].(int) * int(time.Second)),
	}
}

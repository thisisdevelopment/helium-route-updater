package types

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type LnsConfig struct {
	Type        string `env:"TYPE"`
	Listen      string `env:"LISTEN"`
	ApiAuth     string `env:"API_AUTH"`
	ApiEndpoint string `env:"API_ENDPOINT"`
}

type HeliumConfig struct {
	RouteId string `env:"ROUTE_ID"`
	KeyPair string `env:"KEYPAIR"`
	Server  string `env:"SERVER"`
}

type Config struct {
	Helium *HeliumConfig `env:",prefix=HELIUM_"`
	Lns    *LnsConfig    `env:",prefix=LNS_"`
}

func ConfigFromEnv() *Config {
	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	return &c
}

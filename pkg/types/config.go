package types

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
	"os"
)

type LnsConfig struct {
	Type        string `env:"TYPE,required"`
	Listen      string `env:"LISTEN,required"`
	ApiAuth     string `env:"API_AUTH,required"`
	ApiEndpoint string `env:"API_ENDPOINT,required"`
	AutoRoaming bool   `env:"AUTO_ROAMING,default=false"`
}

type HeliumConfig struct {
	RouteId string `env:"ROUTE_ID,required"`
	KeyPair string `env:"KEYPAIR,required"`
	Server  string `env:"SERVER,required"`
}

type Config struct {
	Helium *HeliumConfig `env:",prefix=HELIUM_"`
	Lns    *LnsConfig    `env:",prefix=LNS_"`
}

func ConfigFromEnv() *Config {
	ctx := context.Background()
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	return &c
}

package urlshortener

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GatewayServerAddr  string
	ShortenServerAddr  string
	RetrieveServerAddr string
}

var AppCfg Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	AppCfg = Config{
		GatewayServerAddr:  getEnvVar("GATEWAY_SERVER_ADDR"),
		ShortenServerAddr:  getEnvVar("SHORTEN_SERVER_ADDR"),
		RetrieveServerAddr: getEnvVar("RETRIEVE_SERVER_ADDR"),
	}

}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("required environment variable %s not set", key)
	}
	return val
}

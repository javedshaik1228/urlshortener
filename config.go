package urlshortener

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	GatewayServerAddr  string
	ShortenServerAddr  string
	RetrieveServerAddr string
}

type DbConfig struct {
	DbUser     string
	DbPwd      string
	Database   string
	Collection string
	Host       string
}

var (
	AppCfg   AppConfig
	DbCfg    DbConfig
	DbCnxUri string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	AppCfg = AppConfig{
		GatewayServerAddr:  getEnvVar("GATEWAY_SERVER_ADDR"),
		ShortenServerAddr:  getEnvVar("SHORTEN_SERVER_ADDR"),
		RetrieveServerAddr: getEnvVar("RETRIEVE_SERVER_ADDR"),
	}

	DbCfg = DbConfig{
		DbUser:     getEnvVar("DB_USER"),
		DbPwd:      getEnvVar("DB_PASSWORD"),
		Collection: getEnvVar("COLLECTION_NAME"),
		Database:   getEnvVar("DATABASE_NAME"),
		Host:       getEnvVar("HOST"),
	}

	DbCnxUri = constructDbUri(DbCfg)
}

func constructDbUri(cfg DbConfig) string {
	return fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		cfg.DbUser, cfg.DbPwd, cfg.Host, cfg.Database,
	)
}

func getEnvVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("required environment variable %s not set", key)
	}
	return val
}

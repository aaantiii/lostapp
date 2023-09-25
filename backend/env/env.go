package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariable string // EnvironmentVariable type represents an environment variable.

const (
	MODE                       EnvironmentVariable = "MODE"
	DOMAIN                     EnvironmentVariable = "DOMAIN"
	PORT                       EnvironmentVariable = "PORT"
	FRONTEND_URL               EnvironmentVariable = "FRONTEND_URL"
	POSTGRES_URL               EnvironmentVariable = "POSTGRES_URL"
	COC_API_EMAILS             EnvironmentVariable = "COC_API_EMAILS"
	COC_API_PASSWORDS          EnvironmentVariable = "COC_API_PASSWORDS"
	DISCORD_CLIENT_ID          EnvironmentVariable = "DISCORD_CLIENT_ID"
	DISCORD_CLIENT_SECRET      EnvironmentVariable = "DISCORD_CLIENT_SECRET"
	DISCORD_OAUTH_REDIRECT_URL EnvironmentVariable = "DISCORD_OAUTH_REDIRECT_URL"
)

// Value returns the value of the environment variable as string.
func (v EnvironmentVariable) Value() string {
	return os.Getenv(v.Name())
}

// Name returns the name of the environment variable.
func (v EnvironmentVariable) Name() string {
	return string(v)
}

func Init() error {
	// docker sets env variables in PROD
	if MODE.Value() != "PROD" {
		if err := godotenv.Load(); err != nil {
			return err
		}
		log.Println("Environment variables loaded.")
	}

	requiredEnv := []EnvironmentVariable{
		DOMAIN,
		PORT,
		FRONTEND_URL,
		POSTGRES_URL,
		COC_API_EMAILS,
		COC_API_PASSWORDS,
		DISCORD_CLIENT_ID,
		DISCORD_CLIENT_SECRET,
		DISCORD_OAUTH_REDIRECT_URL,
	}

	for _, envVar := range requiredEnv {
		if _, found := os.LookupEnv(envVar.Name()); !found {
			return fmt.Errorf("required env variable '%s' not set", envVar.Name())
		}
	}

	log.Print("All required env variables are set.")
	return nil
}

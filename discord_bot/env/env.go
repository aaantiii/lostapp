package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariable string // EnvironmentVariable type represents an environment variable.

const (
	MODE                  EnvironmentVariable = "MODE"
	POSTGRES_URL          EnvironmentVariable = "POSTGRES_URL"
	DISCORD_CLIENT_ID     EnvironmentVariable = "DISCORD_CLIENT_ID"
	DISCORD_CLIENT_SECRET EnvironmentVariable = "DISCORD_CLIENT_SECRET"
	DISCORD_GUILD_ID      EnvironmentVariable = "DISCORD_GUILD_ID"
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
		if err := godotenv.Load("../.env"); err != nil {
			return err
		}
		log.Println("Environment variables loaded.")
	}

	requiredEnv := []EnvironmentVariable{
		MODE,
		POSTGRES_URL,
		DISCORD_CLIENT_ID,
		DISCORD_CLIENT_SECRET,
		DISCORD_GUILD_ID,
	}

	for _, envVar := range requiredEnv {
		if _, found := os.LookupEnv(envVar.Name()); !found {
			return fmt.Errorf("required env variable '%s' not set", envVar.Name())
		}
	}

	log.Println("All required env variables are set.")
	return nil
}

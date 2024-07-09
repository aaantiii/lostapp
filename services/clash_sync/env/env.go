package env

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariable string // EnvironmentVariable type represents an environment variable.

const (
	MODE         EnvironmentVariable = "MODE"
	POSTGRES_URL EnvironmentVariable = "POSTGRES_URL"
	COC_EMAIL    EnvironmentVariable = "COC_EMAIL"
	COC_PASSWORD EnvironmentVariable = "COC_PASSWORD"
)

// Value returns the value of the environment variable as string.
func (v EnvironmentVariable) Value() string {
	return os.Getenv(v.Name())
}

// Name returns the name of the environment variable.
func (v EnvironmentVariable) Name() string {
	return string(v)
}

func Load() error {
	// in prod mode the env variables are set by systemd service
	if MODE.Value() != "PROD" {
		if err := godotenv.Load("../.env"); err != nil {
			return err
		}
	}

	required := []EnvironmentVariable{
		MODE,
		POSTGRES_URL,
		COC_EMAIL,
		COC_PASSWORD,
	}

	for _, v := range required {
		if _, found := os.LookupEnv(v.Name()); !found {
			return fmt.Errorf("required env variable '%s' not set", v.Name())
		}
	}

	slog.Info("All required env variables are set.")
	return nil
}

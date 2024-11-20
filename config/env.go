package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type environments struct {
	// Config Server HTTP
	GIN_MODE string `validate:"required"`
	PORT     string `validate:"required"`

	// Config Type Env
	GO_ENV string `validate:"required"`
}

func loadEnv() (*environments, error) {
	// Use OS environment variables if config file is not found
	useOS := false

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		useOS = true
	}

	// Load environment variables
	env := &environments{
		GIN_MODE: getEnv("GIN_MODE", useOS),
		PORT:    getEnv("PORT", useOS),

		GO_ENV:  getEnv("GO_ENV", useOS),
	}

	// Validate struct
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(env); err != nil {
		return nil, err
	}

	return env, nil
}

func getEnv(key string, useOS bool) string {
	if useOS {
		return os.Getenv(key)
	}
	
	return viper.GetString(key)
}
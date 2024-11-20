package config

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Container struct {
	Environments *environments
	Logger       *logrus.Logger
}

type applicationMode string
type ginMode string

const aplicationModeKey applicationMode = "aplicationMode"
const ginModeKey ginMode = "ginMode"

func LoadContainer() (context.Context, *Container, error) {
	// Carregar contexto
	ctx := context.Background()

	// Carrega modo de aplicação no contexto
	ctx = context.WithValue(ctx, aplicationModeKey, applicationMode("release"))

	// Carregar logger
	logger := loadLogger(ctx)

	// Carregar configurações do ambiente
	environments, err := loadEnv()

	if err != nil {
		return nil, nil, err
	}

	// Carrega mode do gin no contexto
	ctx = context.WithValue(ctx, ginModeKey, ginMode(environments.GIN_MODE))

	if err != nil {
		return nil, nil, err
	}

	return ctx, &Container{
		Environments: environments,
		Logger:       logger,
	}, nil
}

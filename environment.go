package serializer

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

const (
	GroupDir = "GROUP_DIR"
)

type EnvironmentContract interface {
	Load(envFilePath string) (Environment, error)
	Get(key string) string
	GetGroupDir() string
}

type Environment struct {
	isLoaded bool
}

func (e *Environment) Load(envFilePath string) (Environment, error) {
	if err := godotenv.Load(envFilePath); err != nil {
		return *e, errors.New("Error loading .env file")
	}

	e.isLoaded = true

	return *e, nil
}

func (e *Environment) Get(key string) string {
	if !e.isLoaded {
		return ""
	}

	return os.Getenv(key)
}

func (e *Environment) GetGroupDir() string {
	if !e.isLoaded {
		return ""
	}

	return os.Getenv(GroupDir)
}

package config

import (
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv(env string) {
	if strings.Contains(env, "/") || strings.Contains(env, `\`) {
		godotenv.Load(env)
	} else {
		godotenv.Load("./configs/" + env)
	}
}

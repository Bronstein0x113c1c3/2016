package common

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadToken(dir string) (string, error) {
	err := godotenv.Load(dir)
	if err != nil {
		return "", err
	}
	return os.Getenv("TOKEN"), nil
}

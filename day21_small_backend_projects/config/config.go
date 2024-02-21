package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() (host string, port int, username string, password string, dbname string, db_port int) {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	host = os.Getenv("HOST")
	db_port, _ = strconv.Atoi(os.Getenv("DB"))
	port, _ = strconv.Atoi(os.Getenv("PORT"))
	username = os.Getenv("DB_ACC")
	password = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB_NAME")
	return
}

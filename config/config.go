package config

import (
	"log"
	"os"
)

type Config struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

func GetConfigs() *Config {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DATABASE")
	log.Printf("configs %s, %s, %s, %s ", host, user, pass, db)
	return &Config{
		Host:   host,
		User:   user,
		Pass:   pass,
		DBName: db,
	}

}

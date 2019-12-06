package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Host     string
	User     string
	Pass     string
	DbName   string
	Port     string
	Protocol string
	DbArgs   string
}

func GetConfigs() *Config {

	host := getParamString("MYSQL_DB_HOST", "172.0.0.1")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "")
	dbname := getParamString("MYSQL_DB", "test_db")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	return &Config{
		Host:     host,
		User:     user,
		Pass:     pass,
		DbName:   dbname,
		Port:     port,
		Protocol: protocol,
		DbArgs:   dbargs,
	}
}

func GetConnectionString(config *Config) string {

	if strings.Trim(config.DbArgs, " ") != "" {
		config.DbArgs = "?" + config.DbArgs
	} else {
		config.DbArgs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		config.User, config.Pass, config.Protocol, config.Host, config.Pass, config.DbName, config.DbArgs)
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	log.Printf("param= %s value= %s, deafultvalue= %s", param, env, defaultValue)
	if env != "" {
		return env
	}
	return defaultValue
}

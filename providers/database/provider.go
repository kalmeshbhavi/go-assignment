package database

import (
	"database/sql"
	"fmt"
	"github.com/kalmeshbhavi/go-assignment/config"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseProvider interface {
	GetKnightRepository() KnightRepository
}

type Provider struct {
	DB *sql.DB
}

func (provider *Provider) GetKnightRepository() KnightRepository {
	return &knightRepository{provider: provider}
}

func (provider *Provider) Close() {
	err := provider.DB.Close()
	if err != nil {
		log.Println(err)
	}
}

func NewProvider(connString string) *Provider {
	log.Printf("connection string : %s", connString)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{DB: db}
}

func GetConnectionString(config *config.Config) string {

	if strings.Trim(config.DbArgs, " ") != "" {
		config.DbArgs = "?" + config.DbArgs
	} else {
		config.DbArgs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		config.User, config.Pass, config.Protocol, config.Host, config.Pass, config.DbName, config.DbArgs)
}

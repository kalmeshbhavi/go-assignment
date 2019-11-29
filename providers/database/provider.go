package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kalmeshbhavi/go-assignment/engine"
)

type Provider struct {
	DB *sql.DB
}

func (provider *Provider) GetKnightRepository() engine.KnightRepository {
	return &knightRepository{provider: provider}
}

func (provider *Provider) Close() {
	provider.DB.Close()
}

func NewProvider(connString string) *Provider {
	// root:Welcome123@tcp(localhost:3306)/inventory?parseTime=true"
	//connectionString :=
	//	fmt.Sprintf("user=%s password=%s dbname=%s", "root", "Welcome123", "test")

	var err error
	//db, err := sql.Open("mysql", "root:Welcome123@tcp(localhost:3306)/inventory")
	log.Printf("connection string : %s", connString)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{DB: db}
}

func GetConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "172.0.0.1")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "")
	dbname := getParamString("MYSQL_DB", "test_db")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	log.Printf("param= %s value= %s, deafultvalue= %s", param, env, defaultValue)
	if env != "" {
		return env
	}
	return defaultValue
}

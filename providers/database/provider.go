package database

import (
	"database/sql"
	"log"

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

func NewProvider() *Provider {
	// root:Welcome123@tcp(localhost:3306)/inventory?parseTime=true"
	//connectionString :=
	//	fmt.Sprintf("user=%s password=%s dbname=%s", "root", "Welcome123", "test")

	var err error
	db, err := sql.Open("mysql", "root:Welcome123@tcp(localhost:3306)/inventory")
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{DB: db}
}

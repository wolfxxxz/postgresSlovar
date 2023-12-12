package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DBconnect string
	DB        *sqlx.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

var Conf Postgres

var (
	path = "config/.env"
)

func init() {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("could not find .env file: ", err)
	}

	fmt.Println("PATH", path)
	Conf.DBconnect = os.Getenv("path")
	fmt.Println(Conf.DBconnect)

	Conf.DB, err = sqlx.Connect("postgres", Conf.DBconnect)
	if err != nil {
		log.Fatal(err)
	}

}

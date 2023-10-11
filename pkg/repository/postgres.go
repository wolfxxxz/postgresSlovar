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
	// Достать путь к постгрес из файла
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("could not find .env file: ", err)
	}
	//conf := repository.NewPostgres()
	fmt.Println("PATH", path)
	Conf.DBconnect = os.Getenv("path")
	fmt.Println(Conf.DBconnect)

	// Проверить соединение
	// Connect - сам проверяет соединение
	// Ping не нужен
	Conf.DB, err = sqlx.Connect("postgres", Conf.DBconnect)
	if err != nil {
		log.Fatal(err)
	}

}

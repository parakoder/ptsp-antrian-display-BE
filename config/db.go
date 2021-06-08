package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sqlx.DB
}

var dbConn = &DB{}

func ConnectSQL() (*DB, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("PASS")
	hostName := os.Getenv("HOST")
	userName := os.Getenv("USER_DB")
	hostPort := os.Getenv("PORT")

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		userName,
		password,
		hostName,
		hostPort,
		dbName)
		// log.Println("DB COnn ", d)
		d, err := sqlx.Open("postgres", url)
		if err != nil {
			panic(err)
		}
		d.SetMaxIdleConns(10)
		d.SetMaxOpenConns(10)
		
		dbConn.SQL = d
		return dbConn, err
}
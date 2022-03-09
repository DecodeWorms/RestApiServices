package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"restapiservices/config"

	_ "github.com/lib/pq"
)

type Conn struct {
	Client *sql.DB
}

func NewConn(ctx context.Context, cf config.Config) *Conn {
	// var c *sql.DB
	log.Println("Connecting to DB..")
	dsn := fmt.Sprintf("host=%s dbname=%s port=%s user=%s sslmode=disable", cf.DatabaseHost, cf.DatabaseName, cf.DatabasePort, cf.DatabaseUserName)
	c, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB.")

	db := Conn{
		Client: c,
	}
	return &db

}

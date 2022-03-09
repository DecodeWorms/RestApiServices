package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"restapiservices/config"
	"restapiservices/database"
	"restapiservices/handler"
	serverjson "restapiservices/serverJson"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var conn *database.Conn
var auth *database.AuthorServices
var hnd *handler.AuthorHandler

//var js *serverjson.AuthorJson

func init() {

	_ = godotenv.Load()
	h := os.Getenv("LOCALHOST")
	db := os.Getenv("DATABASE_NAME")
	p := os.Getenv("DATABASE_PORT")
	u := os.Getenv("DATABASE_USERNAME")

	cfg := config.Config{
		DatabaseHost:     h,
		DatabaseName:     db,
		DatabasePort:     p,
		DatabaseUserName: u,
	}
	var ctx context.Context
	conn = database.NewConn(ctx, cfg)
	auth = database.NewAuthorServices(conn)
	hnd = handler.NewAuthorHandler(auth)

}

func main() {
	router := mux.NewRouter()
	ser := serverjson.NewAuthJosn(hnd)
	router.HandleFunc("/auth/create", ser.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/auth/authors", ser.Authors).Methods(http.MethodGet)
	router.HandleFunc("/auth/author/{email}", ser.Author).Methods(http.MethodGet)
	router.HandleFunc("/auth/update", ser.Update).Methods(http.MethodPut)
	router.HandleFunc("/auth/del/{email}", ser.Delete).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":9000", router))

}

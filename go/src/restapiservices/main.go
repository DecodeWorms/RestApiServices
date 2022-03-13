package main

import (
	"context"
	"os"
	"restapiservices/config"
	"restapiservices/database"
	"restapiservices/handler"
	serverjson "restapiservices/serverJson"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	ser := serverjson.NewAuthJosn(hnd)
	router.POST("auth/create", ser.SignUp())
	router.GET("auth/author/:email", ser.Author())
	router.GET("/auth/authors", ser.Authors())
	router.PUT("/auth/update", ser.Update())
	router.DELETE("/auth/delete/:email", ser.Delete())
	// router.HandleFunc("/auth/authors", ser.Authors).Methods(http.MethodGet)
	// router.HandleFunc("/auth/author/{email}", ser.Author).Methods(http.MethodGet)
	// router.HandleFunc("/auth/update", ser.Update).Methods(http.MethodPut)
	// router.HandleFunc("/auth/del/{email}", ser.Delete).Methods(http.MethodDelete)
	// router.HandleFunc("/auth/valem", ser.ValidateEmail).Methods(http.MethodGet)
	router.Run(":9000")
	//log.Fatal(http.ListenAndServe(":9000", router))

}

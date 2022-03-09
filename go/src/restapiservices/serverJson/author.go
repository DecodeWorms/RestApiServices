package serverjson

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"restapiservices/handler"
	"restapiservices/objects"
	"restapiservices/util"

	"github.com/gorilla/mux"
)

type AuthorJson struct {
	author *handler.AuthorHandler
}

func NewAuthJosn(author *handler.AuthorHandler) *AuthorJson {
	au := &AuthorJson{
		author: author,
	}
	return au

}

func (au AuthorJson) SignUp(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var ctx context.Context
	var auths objects.Author
	if err := json.NewDecoder(r.Body).Decode(&auths); err != nil {
		json.NewEncoder(w).Encode("Unable to UnMarshal JSON")
	}
	res, err := au.author.SignUp(ctx, auths)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to create author")
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(res)

}

func (au AuthorJson) Author(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	params := mux.Vars(r)
	email := params["email"]
	var ctx context.Context
	res, err := au.author.Author(ctx, email)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)
}

func (au AuthorJson) Authors(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var ctx context.Context
	res, err := au.author.Authors(ctx)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(res)
}

func (au AuthorJson) Update(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var ctx context.Context
	var author *objects.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := au.author.Update(ctx, author); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode("Username name updated successfully")
}

func (au AuthorJson) Delete(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	params := mux.Vars(r)
	email := params["email"]

	var ctx context.Context
	if err := au.author.Delete(ctx, email); err != nil {
		json.NewEncoder(w).Encode("error in deleteing author")
	}
	json.NewEncoder(w).Encode("author deleted successfully..")
}

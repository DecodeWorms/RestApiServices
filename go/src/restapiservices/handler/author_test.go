package handler

import (
	"fmt"
	"log"
	"restapiservices/objects"
	"testing"
)

func TestValidateRegDatas(t *testing.T) {
	var err error

	datas := objects.Author{Id: 1, Name: "Joe", Email: "joeman@gmail.com", Gender: "male"}

	_, err = ValidateRegDatas(datas)
	if err != nil {
		t.Error(err)

	}
}

func TestValidateAuthor(t *testing.T) {

	values := &objects.Author{Id: 1, Name: "ham", Email: "email", Gender: "male"}
	var err error
	res, err := ValidateAuthor(values)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)

}

func TestValidateEmail(t *testing.T) {
	var name = "name"

	res, err := ValidateEmail(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func TestValidateUpdate(t *testing.T) {

	par := &objects.Author{Email: "Biola@gmail.com", Name: ""}
	res, err := ValidateUpdate(par)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}

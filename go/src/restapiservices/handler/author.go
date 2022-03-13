package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"restapiservices/objects"
)

type AuthorHandler struct {
	authorService objects.AuthorService
}

func NewAuthorHandler(authService objects.AuthorService) *AuthorHandler {
	auth := &AuthorHandler{
		authorService: authService,
	}
	return auth
}

func (au AuthorHandler) SignUp(ctx context.Context, author objects.Author) (*objects.Author, error) {
	res, err := ValidateRegDatas(author)
	if err != nil {
		log.Fatal(err)
	}

	at := &objects.Author{
		Id:     res.Id,
		Email:  res.Email,
		Name:   res.Name,
		Gender: res.Gender,
	}

	if err := au.authorService.Create(ctx, at); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error in creating author")

	}
	return at, nil
}

func (au AuthorHandler) Authors(ctx context.Context) ([]*objects.Author, error) {

	authors, err := au.authorService.Authors(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error in fetching authors %v", err)
	}
	return authors, nil
}

func (au AuthorHandler) Author(ctx context.Context, email string) (*objects.Author, error) {
	e, err := ValidateEmail(email)
	if err != nil {
		log.Fatal(err)
	}
	author, err := au.authorService.Author(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("error in fetching auhor %v", err)
	}

	res, err := ValidateAuthor(author)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (au AuthorHandler) Update(ctx context.Context, at *objects.Author) error {
	res, err := ValidateUpdate(at)
	if err != nil {
		log.Fatal(err)
	}

	if err := au.authorService.Update(ctx, res); err != nil {
		return err
	}
	return nil

}

func (au AuthorHandler) Delete(ctx context.Context, email string) error {
	res, err := ValidateEmail(email)
	if err != nil {
		log.Fatal(err)
	}

	if err := au.authorService.Delete(ctx, res); err != nil {
		return err
	}
	return nil

}

func (au AuthorHandler) ValidateEmail(ctx context.Context, email string) (string, error) {
	em, err := ValidateEmail(email)
	if err != nil {
		log.Fatal(err)
	}

	ma, err := au.authorService.ValidateEmail(ctx, em)
	if err != nil {
		return "", err
	}
	return ma, nil

}

func ValidateRegDatas(data objects.Author) (*objects.Author, error) {

	if data.Id == 0 {
		return nil, errors.New("an id is empty")
	}
	if data.Name == "" {
		return nil, errors.New("name is empty")
	}

	if data.Email == "" {
		return nil, errors.New("an email is empty")
	}

	if data.Gender == "" {
		return nil, errors.New("gender is empty")
	}
	return &data, nil
}

func ValidateAuthor(res *objects.Author) (*objects.Author, error) {

	if res.Id == 0 {
		return nil, errors.New("value an id is empty")
	}
	if res.Name == "" {
		return nil, errors.New("value for name is empty")
	}

	if res.Gender == "" {
		return nil, errors.New("value for gender is empty")
	}
	return res, nil
}

func ValidateEmail(email string) (string, error) {
	if email == "" {
		return "", errors.New("email is empty")
	}
	return email, nil
}

func ValidateUpdate(u *objects.Author) (*objects.Author, error) {

	if u.Email == "" {
		return nil, fmt.Errorf("an empty email supplied")
	}

	if u.Name == "" {
		return nil, fmt.Errorf("an empty name supplied")
	}
	return u, nil
}

// func CheckIfAExist(e string)(string,error){
// 	if e == ""{
// 		return "", errors.New("")
// 	}
// }

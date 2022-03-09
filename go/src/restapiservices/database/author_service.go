package database

import (
	"context"
	"log"
	"restapiservices/objects"
)

type AuthorServices struct {
	conn *Conn
}

func NewAuthorServices(c *Conn) *AuthorServices {
	au := AuthorServices{
		conn: c,
	}
	return &au
}
func (au AuthorServices) Create(ctx context.Context, cr *objects.Author) error {

	c := &objects.Author{
		Id:     cr.Id,
		Email:  cr.Email,
		Name:   cr.Name,
		Gender: cr.Gender,
	}

	_, err := au.conn.Client.Exec("INSERT INTO authors VALUES($1,$2,$3,$4)", c.Id, c.Email, c.Name, c.Gender)
	if err != nil {
		return err
	}
	return nil

}

func (au AuthorServices) Authors(ctx context.Context) ([]*objects.Author, error) {
	var auth []*objects.Author
	rows, err := au.conn.Client.Query("SELECT * FROM authors")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		aut := &objects.Author{}
		if err := rows.Scan(&aut.Id, &aut.Email, &aut.Name, &aut.Gender); err != nil {
			return nil, err
		}
		auth = append(auth, aut)
	}
	return auth, nil

}

func (au AuthorServices) Author(ctx context.Context, email string) (*objects.Author, error) {
	auth := &objects.Author{}
	r := au.conn.Client.QueryRow("SELECT * FROM authors WHERE email = $1", email)

	if err := r.Scan(&auth.Id, &auth.Email, &auth.Name, &auth.Gender); err != nil {
		return nil, err
	}
	return auth, nil

}

func (au AuthorServices) Update(ctx context.Context, author *objects.Author) error {
	_, err := au.conn.Client.Exec("UPDATE authors SET name = $1 WHERE email = $2", author.Name, author.Email)
	if err != nil {
		return nil
	}
	return err

}

func (au AuthorServices) Delete(ctx context.Context, email string) error {
	_, err := au.conn.Client.Exec("DELETE FROM authors WHERE email = $1", email)
	if err != nil {
		return nil
	}
	return err

}

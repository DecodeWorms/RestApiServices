package serverjson

import (
	"context"
	"encoding/json"
	"net/http"
	"restapiservices/handler"
	"restapiservices/objects"
	"restapiservices/util"

	"github.com/gin-gonic/gin"
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

func (au AuthorJson) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var auths objects.Author
		if err := c.ShouldBindJSON(&auths); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := au.author.SignUp(c, auths)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error in creating user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "user created successfully", "results": auths})

	}
}

func (au AuthorJson) Author() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Param("email")
		res, err := au.author.Author(c, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in fetching an author"})
		}
		c.JSON(http.StatusOK, gin.H{"status": res})

	}

}

func (au AuthorJson) Authors() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := au.author.Authors(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in fetching authors"})
		}
		c.JSON(http.StatusOK, gin.H{"results": res})

	}

}

func (au AuthorJson) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var author *objects.Author
		if err := c.ShouldBindJSON(&author); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in Unmarshaling JSON"})
			return
		}
		if err := au.author.Update(c, author); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in updating an author"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "successfully updated"})

	}

}

func (au AuthorJson) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		if err := au.author.Delete(c, email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in Unmarshaling JSON"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "author successfully deleted"})

	}

}

func (au AuthorJson) ValidateEmail(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var ob *objects.Author
	if err := json.NewDecoder(r.Body).Decode(&ob); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	var ctx context.Context
	mail, err := au.author.ValidateEmail(ctx, ob.Email)
	if err != nil {
		json.NewEncoder(w).Encode("error in validating emaill")
	}
	json.NewEncoder(w).Encode(mail)
}

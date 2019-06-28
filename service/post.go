package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-todo/util"
	"go-todo/database"
	"go-todo/model"
)

// PostHandler ...
func PostHandler(c *gin.Context) {
	db := database.Connect(c)
	defer db.Close()

	t := model.Todo{}
	err := c.ShouldBindJSON(&t)
	util.BadRequest(c, err)

	query := "INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id"
	row := db.QueryRow(query, t.Title, t.Status)
	err = row.Scan(&t.ID)
	util.InternalServerError(c, err)
	c.JSON(http.StatusCreated, t)
}
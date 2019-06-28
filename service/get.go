package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"go-todo/util"
	"go-todo/database"
	"go-todo/model"
)

// GetHandler ...
func GetHandler(c *gin.Context) {
	db := database.Connect(c)
	defer db.Close()

	query := "SELECT id, title, status FROM todos"
	stmt, err := db.Prepare(query)
	util.InternalServerError(c, err)
	
	rows, err := stmt.Query()
	util.InternalServerError(c, err)

	todos := []model.Todo{}

	for rows.Next() {
		t := model.Todo{}
		err = rows.Scan(&t.ID, &t.Title, &t.Status)
		util.InternalServerError(c, err)
		todos = append(todos, t)
		fmt.Println("one row ", t.ID, t.Title, t.Status)
	}

	c.JSON(http.StatusOK, todos)
}

// GetByIDHandler ...
func GetByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	util.BadRequest(c, err)

	db := database.Connect(c)
	defer db.Close()

	query := "SELECT id, title, status FROM todos WHERE id=$1;"
	stmt, err := db.Prepare(query)
	util.InternalServerError(c, err)
	
	t := model.Todo{}
	row := stmt.QueryRow(id)
	err = row.Scan(&t.ID, &t.Title, &t.Status)
	util.InternalServerError(c, err)

	c.JSON(http.StatusOK, t)
}
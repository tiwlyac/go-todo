package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"go-todo/database"
	"go-todo/util"
	"go-todo/model"
)

// PutByIDHandler ...
func PutByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	util.BadRequest(c, err)

	db := database.Connect(c)
	defer db.Close()

	t := model.Todo{}
	err = c.ShouldBindJSON(&t)
	util.BadRequest(c, err)
	t.ID = id

	query := "UPDATE todos SET title = $1, status = $2 WHERE id=$3;"
	stmt, err := db.Prepare(query)
	util.InternalServerError(c, err)

	stmt.Exec(t.Title, t.Status, t.ID)
	c.JSON(http.StatusOK, t)
}
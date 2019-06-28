package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"go-todo/util"
	"go-todo/database"
)

// DeleteByIDHandler ...
func DeleteByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	util.BadRequest(c, err)

	db := database.Connect(c)
	defer db.Close()

	query := "DELETE FROM todos WHERE id=$1;"
	stmt, err := db.Prepare(query)
	util.InternalServerError(c, err)

	stmt.Exec(id)
	c.JSON(http.StatusOK, gin.H{ "status": "success" })
}
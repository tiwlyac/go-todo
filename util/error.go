package util

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}

func BadRequest(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
package database

import (
	"os"
	"log"
	"fmt"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" 
	"go-todo/util"
)

// Connect ...
func Connect(c *gin.Context) *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url) 
	util.InternalServerError(c, err)
	return db
}

// Create ...
func Create() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("fatal", err.Error())
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS todos(
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	)
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table", err.Error())
	}  

	fmt.Println("Okay")
}


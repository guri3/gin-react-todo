package main

import (
	"database/sql"
	"net/http"

	"github.com/guri3/todo/src/api/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gin_todo_dev")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	tctr := &controller.Todo{DB: db}
	r.GET("/todos", tctr.All)
	r.POST("/todos", tctr.Create)
	r.POST("/todos/:id/toggle", tctr.ToggleDone)

	r.Run(":3030")
}

package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/guri3/todo/src/api/httputil"
	"github.com/guri3/todo/src/api/model"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	DB *sql.DB
}

func (t *Todo) All(c *gin.Context) {
	todos, err := model.TodoAll(t.DB)

	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")

	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(todos) == 0 {
		c.JSON(http.StatusOK, make([]*model.Todo, 0))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": todos,
		"error":  nil,
	})
}

func (t *Todo) Create(c *gin.Context) {
	var todo model.Todo

	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")

	if c.Request.ContentLength == 0 {
		resp := httputil.NewErrorResponse(errors.New("body is missing"))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := c.BindJSON(&todo); err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	inserted, err := todo.Insert(t.DB)
	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": inserted,
		"error":  nil,
	})
}

func (t *Todo) ToggleDone(c *gin.Context) {
	todo, err := model.TodoByID(t.DB, c.Param("id"))

	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")

	toggled, err := todo.Toggle(t.DB)
	if err != nil {
		resp := httputil.NewErrorResponse(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": toggled,
		"error":  nil,
	})
}

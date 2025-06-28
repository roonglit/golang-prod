package controllers

import (
	"learning/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) CreateTodo(c *gin.Context) {
	var todo Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdTodo, err := s.Store.CreateTodo(
		c,
		createTodoParams(todo),
	)

	c.JSON(http.StatusCreated, todoResponse(createdTodo))
}

func createTodoParams(todo Todo) models.CreateTodoParams {
	return models.CreateTodoParams{
		Title:       todo.Title,
		Description: pgtype.Text{String: todo.Description, Valid: true},
	}
}

func todoResponse(todo models.Todo) Todo {
	return Todo{
		ID:          int(todo.ID),
		Title:       todo.Title,
		Description: todo.Description.String,
		CreatedAt:   todo.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:   todo.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}

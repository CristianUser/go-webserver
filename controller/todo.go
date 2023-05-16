package controller

import (
	"log"
	"net/http"
	"pronesoft/server/model"

	"github.com/gin-gonic/gin"
)

type NewTodo struct {
	Name string `json:"name" binding:"required"`
	Done bool   `json:"done" binding:"required"`
}

type TodoUpdate struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func GetTodos(c *gin.Context) {

	var todos []model.Todo

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&todos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)

}

func GetTodo(c *gin.Context) {

	var todo model.Todo

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)

}

func PostTodo(c *gin.Context) {

	var todo NewTodo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo := model.Todo{Name: todo.Name, Done: todo.Done}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTodo)
}

func UpdateTodo(c *gin.Context) {

	var todo model.Todo

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
		return
	}

	var updateTodo TodoUpdate

	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&todo).Updates(model.Todo{Name: updateTodo.Name, Done: updateTodo.Done}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)

}

func DeleteTodo(c *gin.Context) {

	var todo model.Todo

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})

}

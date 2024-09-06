package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unknownn17/Internship_Task/internal/interface/impliment"
	"github.com/unknownn17/Internship_Task/internal/models"
)

type Handler struct {
	S *impliment.Service
	C context.Context
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.Register_User true "Register User"
// @Success 200 {object} string "Register"
// @Failure 500 {object} string "Error message"
// @Router /user/register [post]
func (u *Handler) Register(c *gin.Context) {
	fmt.Println("hey request came")
	var req models.Register_User

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(500, err)
		return
	}
	res, err := u.S.Register(u.C, &req)
	if err != nil {
		c.JSON(500, err)
	}
	fmt.Printf("it's response %v and so",res)
	c.JSON(200, res)
}

// Verify godoc
// @Summary Verify a user account
// @Description Verify user registration with a verification code
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.Verify_User true "Verify User"
// @Success 200 {object} string "Verify"
// @Failure 500 {object} string "Error message"
// @Router /user/verify [post]
func (u *Handler) Verify(c *gin.Context) {
	var req models.Verify_User
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(500, err)
		return
	}
	res, err := u.S.Verify(u.C, &req)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, res)
}

// LogIn godoc
// @Summary Log in a user
// @Description Authenticate a user with email and password
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.LogIn true "LogIn User"
// @Success 201 {object} string "Login"
// @Failure 500 {object} string "Error message"
// @Router /user/login [post]
func (u *Handler) LogIn(c *gin.Context) {
	var req models.LogIn
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(500, err)
		return
	}
	res, err := u.S.LogIn(u.C, &req)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, res)
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a task for a user
// @Tags Task
// @Accept json
// @Produce json
// @Param request body models.Task true "Create Task"
// @Success 201 {object} models.GetTaskResponse
// @Failure 500 {object} string "Error message"
// @Security ApiKeyAuth
// @Router /task [post]
func (u *Handler) CreateTask(c *gin.Context) {
	var req models.Task
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(500, err)
		return
	}
	res, err := u.S.CreateTask(u.C, &req)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, res)
}

// GetTask godoc
// @Summary Get a specific task
// @Description Retrieve a task by ID and user ID
// @Tags Task
// @Produce json
// @Param id query int true "Task ID"
// @Param user_id query int true "User ID"
// @Success 200 {object} models.GetTaskResponse
// @Failure 404 {object} string "Task not found"
// @Failure 500 {object} string "Error message"
// @Security ApiKeyAuth
// @Router /task [get]
func (u *Handler) GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(err)
	}
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println(err)
	}

	res, err := u.S.GetTask(u.C, &models.GetTaskRequest{ID: id, UserID: userid})
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, res)
}

// GetTasks godoc
// @Summary Get all tasks for a user
// @Description Retrieve all tasks for a specific user
// @Tags Task
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {object} []models.GetTaskResponse
// @Failure 500 {object} string "Error message"
// @Security ApiKeyAuth
// @Router /tasks [get]
func (u *Handler) GetsTasks(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println(err)
	}
	res, err := u.S.GetTasks(u.C, userid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, res)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update task information for a specific task and user
// @Tags Task
// @Accept json
// @Produce json
// @Param id query int true "Task ID"
// @Param user_id query int true "User ID"
// @Param request body models.Task true "Update Task"
// @Success 200 {object} models.GetTaskResponse
// @Failure 500 {object} string "Error message"
// @Security ApiKeyAuth
// @Router /task [put]
func (u *Handler) UpdateTask(c *gin.Context) {
	var req models.Task

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(err)
	}
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println(err)
	}
	req.ID = id
	req.UserID = userid
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(500, err)
		return
	}
	res, err := u.S.UpdateTask(u.C, &req)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, res)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID and user ID
// @Tags Task
// @Produce json
// @Param id query int true "Task ID"
// @Param user_id query int true "User ID"
// @Success 200 {object} string "Task deleted"
// @Failure 500 {object} string "Error message"
// @Security ApiKeyAuth
// @Router /task [delete]
func (u *Handler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(err)
	}
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println(err)
	}
	res, err := u.S.DeleteTask(u.C, &models.GetTaskRequest{ID: id, UserID: userid})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, res)
}

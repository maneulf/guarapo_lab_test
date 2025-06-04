package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maneulf/guarapo_lab_test/internal/handlers/models/req"
	"github.com/maneulf/guarapo_lab_test/internal/ports"
)

var tokenLenght = 16
var Authorization = "Authorization"

type Handlers struct {
	*Base
	TasksService ports.TasksService
}

func getToken(ctx *gin.Context) string {
	token := ctx.GetHeader(Authorization)[7:]
	return token
}

func New(base *Base, tasksService ports.TasksService) *Handlers {
	return &Handlers{
		Base:         base,
		TasksService: tasksService,
	}
}

func (h *Handlers) Login(ctx *gin.Context) {
	var username req.Username
	err := ctx.ShouldBindJSON(&username)

	if err != nil {
		log.Printf("Could not read json body, Err: %s", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	token := h.generateToken(tokenLenght)
	log.Print("Token successfuly created")
	h.Usernames[token] = username.Username
	ctx.JSON(200, gin.H{
		"token": token,
	})

}

func (h *Handlers) GetTasks(ctx *gin.Context) {
	token := getToken(ctx)
	tasks, err := h.TasksService.GetTasks(token)

	if err != nil {
		log.Printf("Error trying to get tasks, Err: %s", err)
		ctx.Status(http.StatusNotFound)
		return

	}

	ctx.JSON(http.StatusOK, tasks)
}

func (h *Handlers) GetTask(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Printf("Error trying to get id, Err: %s", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	token := getToken(ctx)
	task, err := h.TasksService.GetTask(taskId, token)

	if err != nil {
		log.Printf("Error trying to get task, Err: %s", err)
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (h *Handlers) SaveTask(ctx *gin.Context) {
	var tasksRequest req.Task
	err := ctx.ShouldBindJSON(&tasksRequest)

	if err != nil {
		log.Printf("Could not read json body, Error: %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "the title field must be filled in"})
		return
	}

	token := getToken(ctx)
	err = h.TasksService.Save(tasksRequest, token)

	if err != nil {
		log.Printf("Error trying to save task, Err: %s", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)

}

func (h *Handlers) UpdateTask(ctx *gin.Context) {
	token := getToken(ctx)
	var tasksRequest req.Task
	err := ctx.ShouldBindJSON(&tasksRequest)

	if err != nil {
		log.Printf("Could not read json body, Error: %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "the title field must be filled in"})
		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("Error trying to get id, Err: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not read id param"})
		return
	}

	err = h.TasksService.Update(tasksRequest, taskId, token)
	if err != nil {
		log.Printf("Error trying to update data, Err: %s", err)
		ctx.Status(http.StatusBadRequest)
		return

	}
	ctx.Status(http.StatusNoContent)

}

func (h *Handlers) DeleteTask(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("Error trying to get id, Err: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not read id param"})
		return
	}

	token := getToken(ctx)

	err = h.TasksService.Delete(taskId, token)
	if err != nil {
		log.Printf("Error trying to delete data, Err: %s", err)
		ctx.Status(http.StatusBadRequest)
		return

	}
	ctx.Status(http.StatusNoContent)
}

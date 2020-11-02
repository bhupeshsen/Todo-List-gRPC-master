package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	cm "github.com/joaquinto/Todo-List-gRPC/client/model"
	"github.com/joaquinto/Todo-List-gRPC/client/response"
	"github.com/joaquinto/Todo-List-gRPC/model"
)

type Client struct {
	ServiceClient model.TodoServiceClient
}

func (c *Client) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading data: %v", err.Error())
		return
	}
	todo := &cm.Todo{}
	json.Unmarshal(data, todo)
	errMessage, notValid := cm.ValidateInput(todo)
	if notValid {
		response.JSON(w, http.StatusBadRequest,
			"Validation Error", errMessage)
		return
	}
	todo.Prepare()
	ctx := context.Background()
	request := &model.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}
	res, err := c.ServiceClient.CreateTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusCreated,
		"Todo saved successfully", todosResponse)
}

func (c *Client) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	request := &model.GetTodos{}
	ctx := context.Background()
	res, err := c.ServiceClient.GetAllTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusOK,
		"Todos fetched successfully", todosResponse)
}

func (c *Client) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	ctx := context.Background()
	request := &model.TodoID{
		Id: todoID,
	}
	res, err := c.ServiceClient.GetTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todoResponse := &model.TodoResponse{
		Todo: res.GetTodo(),
	}
	response.JSON(w, http.StatusOK,
		"Todo fetched successfully", todoResponse)
}

func (c *Client) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	ctx := context.Background()
	request := &model.TodoID{
		Id: todoID,
	}
	res, err := c.ServiceClient.DeleteTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusOK,
		"Todo fetched successfully", todosResponse)
}

func (c *Client) EditTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading data: %v", err.Error())
		return
	}
	todo := &cm.Todo{}
	json.Unmarshal(data, todo)
	errMessage, notValid := cm.ValidateInput(todo)
	if notValid {
		response.JSON(w, http.StatusBadRequest,
			"Validation Error", errMessage)
		return
	}
	todo.Prepare()
	ctx := context.Background()
	request := &model.Todo{
		Id:          todoID,
		Title:       todo.Title,
		Description: todo.Description,
	}
	res, err := c.ServiceClient.EditTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusOK,
		"Todo Edited successfully", todosResponse)
}

func (c *Client) MarkTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["id"]
	ctx := context.Background()
	request := &model.TodoID{
		Id: todoID,
	}
	res, err := c.ServiceClient.MarkTodo(ctx, request)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, "Internal server error", err)
		return
	}
	todosResponse := &model.TodosResponse{
		Todos: res.GetTodos(),
	}
	response.JSON(w, http.StatusOK,
		"Todo checked successfully", todosResponse)
}

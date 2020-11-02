package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/joaquinto/Todo-List-gRPC/client/handler"
	"github.com/joaquinto/Todo-List-gRPC/model"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
		return
	}
	client := model.NewTodoServiceClient(conn)
	sc := handler.Client{ServiceClient: client}
	fmt.Println("Server started on port 8080")
	router := mux.NewRouter()
	router.HandleFunc("/create-todo", sc.CreateTodoHandler).Methods(http.MethodPost)
	router.HandleFunc("/todos", sc.GetTodosHandler).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", sc.GetTodoHandler).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", sc.MarkTodoHandler).Methods(http.MethodPatch)
	router.HandleFunc("/todos/{id}/edit", sc.EditTodoHandler).Methods(http.MethodPatch)
	router.HandleFunc("/todos/{id}", sc.DeleteTodoHandler).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", router))
}

package handler

import (
	srv "api/gen/service"
	"context"
	"log"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type todos struct {
	mutex     sync.Mutex
	data      []*srv.Todo
	idCounter int32
}

func (t *todos) Create(content string) *srv.Todo {
	t.mutex.Lock()
	now := time.Now()
	newTodo := &srv.Todo{
		Id:      t.idCounter,
		Content: content,
		Status:  srv.Todo_WAITING,
		CreatedAt: &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}
	t.idCounter++
	t.mutex.Unlock()
	return newTodo
}

func (t *todos) Add(newTodo *srv.Todo) {
	t.mutex.Lock()
	t.data = append(t.data, newTodo)
	t.idCounter++
	t.mutex.Unlock()
}

type TodoHandler struct {
	todos *todos
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todos: &todos{
			mutex: sync.Mutex{},
			data:  make([]*srv.Todo, 0),
		},
	}
}

func (h *TodoHandler) GetAll(c context.Context, r *srv.GetAllRequest) (*srv.GetAllResponse, error) {
	log.Println("Received GetAll request!")
	response := &srv.GetAllResponse{
		Todos: h.todos.data,
	}
	return response, nil
}

func (h *TodoHandler) Add(c context.Context, r *srv.AddRequest) (*srv.AddResponse, error) {
	log.Println("Received Add request!")
	newTodo := h.todos.Create(r.Content)
	h.todos.Add(newTodo)

	return &srv.AddResponse{
		Todo: newTodo,
	}, nil
}

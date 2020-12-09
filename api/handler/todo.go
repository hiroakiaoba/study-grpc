package handler

import (
	srv "api/gen/service"
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type todos struct {
	mutex     sync.Mutex
	data      []*srv.Todo
	idCounter int32
}

func (t *todos) Create(content string) *srv.Todo {
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
	t.mutex.Lock()
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
	initialTodos := &todos{
		mutex:     sync.Mutex{},
		data:      []*srv.Todo{},
		idCounter: 1,
	}
	initialTodos.Add(initialTodos.Create("hogehog"))
	return &TodoHandler{
		todos: initialTodos,
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

	// request data check
	content := r.Content
	if content == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Todoの内容を送信してください")
	}

	newTodo := h.todos.Create(content)
	h.todos.Add(newTodo)

	return &srv.AddResponse{
		Todo: newTodo,
	}, nil
}

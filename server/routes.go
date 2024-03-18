package server

import (
	"context"
	"net/http"
	database "nicklatch/fe-mentor-todo-app/database/sqlc"
	"nicklatch/fe-mentor-todo-app/web"
	"strconv"
	"strings"
)

type updateData struct {
	completed string
	id        string
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(web.Files))

	// serves embedded files in the web/assets dir
	mux.Handle("GET /assets/*", fileServer)

	mux.HandleFunc("GET /", s.rootHandler)
	mux.HandleFunc("GET /all", s.getAll)
	mux.HandleFunc("GET /active", s.getAllActive)
	mux.HandleFunc("GET /completed", s.getAllCompleted)

	mux.HandleFunc("POST /", s.newTodoHandler)
	mux.HandleFunc("PUT /{id}", s.updateHandler)
	mux.HandleFunc("DELETE /{id}", s.deleteHandler)

	return mux
}

func (s *Server) rootHandler(rw http.ResponseWriter, req *http.Request) {
	tasks, err := s.db.GetAllTodos(context.Background())
	s.LogError(err)

	err = s.tmpl.ExecuteTemplate(rw, "base", tasks)
	s.LogError(err)
}

func (s *Server) getAll(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllTodos(context.Background())
	s.LogError(err)

	err = s.tmpl.ExecuteTemplate(rw, "todo-list", data)
	s.LogError(err)
}

func (s *Server) getAllActive(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllActiveTodos(context.Background())
	s.LogError(err)

	err = s.tmpl.ExecuteTemplate(rw, "todo-list", data)
	s.LogError(err)
}

func (s *Server) getAllCompleted(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllCompletedTodos(context.Background())
	s.LogError(err)

	err = s.tmpl.ExecuteTemplate(rw, "todo-list", data)
	s.LogError(err)
}

func (s *Server) updateHandler(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.RequestURI(), "/")
	val, err := strconv.Atoi(id)
	s.LogError(err)

	err = s.db.ToggleCompleted(context.Background(), int32(val))
	s.LogError(err)
}

func (s *Server) deleteHandler(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.RequestURI(), "/")
	val, err := strconv.Atoi(id)
	s.LogError(err)

	err = s.db.DeleteOneTodo(context.Background(), int32(val))
	s.LogError(err)
}

func (s *Server) newTodoHandler(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	s.LogError(err)

	b, err := strconv.ParseBool(req.Form.Get("completed"))
	s.LogError(err)

	todo, err := s.db.CreateTodo(context.Background(), database.CreateTodoParams{
		Task:      req.Form.Get("task"),
		Completed: b,
	})
	s.LogError(err)

	err = s.tmpl.ExecuteTemplate(rw, "todo-item", todo)
	s.LogError(err)
}

// LogError is a Server method that takes an error
// and if the error is some, it will log it.
func (s *Server) LogError(err error) {
	if err != nil {
		s.log.Error(err)
	}
}

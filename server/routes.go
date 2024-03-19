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
	mux.HandleFunc("DELETE /many", s.deleteManyHandler)

	return mux
}

func (s *Server) rootHandler(rw http.ResponseWriter, req *http.Request) {
	tasks, err := s.db.GetAllTodos(context.Background())
	if err != nil {
		s.log.Error(err)
	}

	err = s.tmpl.ExecuteTemplate(rw, "base", tasks)
	if err != nil {
		s.log.Error(err)
	}
}

func (s *Server) getAll(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllTodos(context.Background())
	if err != nil {
		s.log.Error(err)
	}

	err = s.tmpl.ExecuteTemplate(rw, "todo-list", data)
	if err != nil {
		s.log.Error(err)
	}
}

func (s *Server) getAllActive(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllActiveTodos(context.Background())
	if err != nil {
		s.log.Error(err)
	}

	err = s.tmpl.ExecuteTemplate(rw, "todo-list", data)
	if err != nil {
		s.log.Error(err)
	}
}

func (s *Server) getAllCompleted(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllCompletedTodos(context.Background())
	if err != nil {
		s.log.Error(err)
	}

	err = s.tmpl.ExecuteTemplate(rw, "todo-list", data)
	if err != nil {
		s.log.Error(err)
	}
}

func (s *Server) updateHandler(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.RequestURI(), "/")
	val, err := strconv.Atoi(id)
	if err != nil {
		s.log.Error(err)
	}

	err = s.db.ToggleCompleted(context.Background(), int32(val))
	if err != nil {
		s.log.Error(err)
	}
}

func (s *Server) deleteHandler(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.RequestURI(), "/")

	val, err := strconv.Atoi(id)
	if err != nil {
		s.log.Error(err)
	}

	err = s.db.DeleteOneTodo(context.Background(), int32(val))
	if err != nil {
		s.log.Error(err)
	}

	rw.WriteHeader(200)
}

func (s *Server) deleteManyHandler(rw http.ResponseWriter, req *http.Request) {
	data, err := s.db.GetAllActiveTodos(context.Background()) //optimistic response for speeeeeeed
	if err != nil {
		s.log.Error(err)
	}

	err = s.db.DeleteCompletedTodos(context.Background())
	if err != nil {
		s.log.Error(err)
	}

	err = s.tmpl.ExecuteTemplate(rw, "todo-list", data)
	if err != nil {
		s.log.Error(err)
	}
}

func (s *Server) newTodoHandler(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		s.log.Error(err)
	}

	b, err := strconv.ParseBool(req.Form.Get("completed"))
	if err != nil {
		s.log.Error(err)
	}

	todo, err := s.db.CreateTodo(context.Background(), database.CreateTodoParams{
		Task:      req.Form.Get("task"),
		Completed: b,
	})
	if err != nil {
		s.log.Error(err)
	}

	err = s.tmpl.ExecuteTemplate(rw, "todo-item", todo)
	if err != nil {
		s.log.Error(err)
	}
}

package server

import (
	"net/http"
	"nicklatch/fe-mentor-todo-app/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(web.Files))

	// serves embedded files in the web/assets dir
	mux.Handle("GET /assets/*", fileServer)
	mux.HandleFunc("GET /", s.rootHandler)

	return mux
}

// rootHandler TODO: serve main page
func (s *Server) rootHandler(rw http.ResponseWriter, req *http.Request) {
	err := s.tmpl.ExecuteTemplate(rw, "base", "")
	if err != nil {
		s.log.Error(err)
	}
	s.log.Infof("%v", req.Method)
}

// TODO: serve all tasks
// TODO: serve one task
// TODO: put one task
// TODO: post one task

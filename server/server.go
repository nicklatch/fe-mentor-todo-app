package server

import (
	"fmt"
	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
	"html/template"
	"net/http"
	database "nicklatch/fe-mentor-todo-app/database/sqlc"
	"os"
	"strconv"
	"time"
)

var files = []string{
	"web/templates/base.gohtml",
}

type Server struct {
	port int
	tmpl *template.Template
	db   *database.Queries
	log  *log.Logger
}

func NewServer(db *database.Queries) *http.Server {
	convPort, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: convPort,
		tmpl: template.Must(template.ParseFiles(files...)),
		db:   db,
		log: log.NewWithOptions(os.Stderr, log.Options{
			ReportCaller:    true,
			ReportTimestamp: true,
			TimeFormat:      time.Kitchen,
		}),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

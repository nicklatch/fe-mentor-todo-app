package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	database "nicklatch/fe-mentor-todo-app/database/sqlc"
	"nicklatch/fe-mentor-todo-app/server"
	"os"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	s := server.NewServer(database.New(conn))

	log.Print("Starting on port :8080")

	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}

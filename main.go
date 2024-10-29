package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
)

func main() {
	connStr := "postgres://postgres:postgres@db.connect.orb.local:5432/postgres?sslmode=disable"
	fmt.Println(connStr)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
		return
	}

	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"github.com/ngtk/practice-apig/db"
	"github.com/ngtk/practice-apig/server"
)

// main ...
func main() {
	database := db.Connect()
	s := server.Setup(database)
	s.Run(":8080")
}

package main

import (
	readinglist "github.com/CASE-fold/gizmo/v2/examples/servers/reading-list"
	"github.com/CASE-fold/gizmo/v2/server/kit"
)

// a tiny main package that simply initializes and initiates the server.
func main() {
	db, err := readinglist.NewDB()
	if err != nil {
		panic(err)
	}
	svc, err := readinglist.NewService(db)
	if err != nil {
		panic(err)
	}
	kit.Run(svc)
}

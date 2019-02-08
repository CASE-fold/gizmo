package main

import (
	readinglist "github.com/case-fold/gizmo/examples/servers/reading-list"
	"github.com/case-fold/gizmo/server/kit"
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

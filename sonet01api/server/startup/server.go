package startup

import (
	"fmt"
	"net/http"
	"sonet01api/server/database"
	// "sonet01api/server/database"
)

func Server(dbpath string, port string) {

	//set up database
	err := create(dbpath, port)
	if err != nil {
		fmt.Println("error starting server", err)
		return
	}

	fmt.Println("Server is running on port 8080")
}

func create(dbpath string, port string) error {

	// initialize sqlite database
	msg := database.InitDB(dbpath)
	if msg != nil {
		fmt.Println("database initialization:", msg)
	}

	RegisterRoutes("v1")

	// initialize server
	server := &http.Server{
		Addr: ":" + port,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("listen: %s\n", err)
	}

	return nil
}

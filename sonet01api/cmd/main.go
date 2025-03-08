package main

import (
	"fmt"
	startup "sonet01api/server/startup"
)

func main() {

	fmt.Println("running app..")

	startup.Server("../sonet01.db", "8080")

}

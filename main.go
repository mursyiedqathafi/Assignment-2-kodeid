package main

import (
	"assignment-2/controller"
	"assignment-2/database"
	"assignment-2/router"
	"fmt"
)

func main() {
	db, err := database.Start()
	if err != nil {
		fmt.Println("Error start db", err)
		return
	}

	ctl := controller.New(db)

	err = router.StartServer(ctl)
	if err != nil {
		fmt.Println("Error start server", err)
		return
	}
}

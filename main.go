package main

import (
	"log"
	"social-dashboard/api/controllers/user"
	"social-dashboard/api/routes"
)

func main() {
	err := user.CreateUsers()
	if err != nil {
		log.Println("create users error : ", err)
		return
	}
	routes.Run()
}

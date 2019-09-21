package main

import "web/router"

func main() {
	router := router.InitRouter()
	router.Run(":8000")
}
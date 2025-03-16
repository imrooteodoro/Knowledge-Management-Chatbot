package main

import (
	"laskerbot/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	botServer := gin.Default()

	routes.SendMessage(botServer)
	routes.InfoBot(botServer)

	botServer.Run(":8080")

}

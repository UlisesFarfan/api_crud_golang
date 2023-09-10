package main

import (
	"os"

	r "github.com/api-rest-go/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	router := gin.Default()

	app := r.Router(router)

	app.Run(port)
}

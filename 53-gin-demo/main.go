package main

import (
	"demo/handlers"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//log.Println("This log message includes date, time, and file information.")
	var (
		PORT string
	)
	PORT = os.Getenv("PORT")

	if PORT == "" {
		flag.StringVar(&PORT, "port", "9091", "--port=9091")
	}
	flag.Parse()

	r := gin.Default() // new instance

	r.Use(handlers.ValidateToken()) // all
	r.GET("/ping", handlers.Ping)

	r.GET("/health", handlers.Health)
	r.POST("/user", handlers.CreateUser)

	r.Run(":" + PORT) // cater to http requests

}

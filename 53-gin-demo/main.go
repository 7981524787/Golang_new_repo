package main

import (
	"demo/database"
	"demo/filedb"
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
		DSN  string
	)
	PORT = os.Getenv("PORT")
	DSN = os.Getenv("DSN")

	if PORT == "" {
		flag.StringVar(&PORT, "port", "9091", "--port=9091")
	}
	flag.Parse()
	if DSN == "" {
		DSN = "host=localhost user=admin password=admin123 dbname=demodb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}

	/// --------------

	/// -------------

	//db, err := database.GetConnection(DSN)
	_, err := database.GetConnection(DSN)
	if err != nil {
		log.Panicln(err)
	}

	r := gin.Default()              // new instance
	r.Use(handlers.ValidateToken()) // all

	//userHandler := handlers.NewUserHandler(database.NewUser(db))
	userHandler := handlers.NewUserHandler(filedb.NewUser("data.txt"))
	r.GET("/ping", handlers.Ping)
	r.GET("/health", handlers.Health)
	r.POST("/user", userHandler.CreateUser())
	r.Run(":" + PORT) // cater to http requests
}

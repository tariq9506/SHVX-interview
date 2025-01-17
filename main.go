package main

import (
	"log"
	"os"
	"shvx/router"

	"github.com/gin-contrib/pprof"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load env")
		return
	}
	r := router.SetupRouter()
	pprof.Register(r)
	r.Run(":" + os.Getenv("PORT"))

}

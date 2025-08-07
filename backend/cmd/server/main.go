package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/example/ludo-game/backend/internal/auth"
	"github.com/example/ludo-game/backend/internal/database"
	"github.com/example/ludo-game/backend/internal/websocket"
)

func main() {
	database.Connect(os.Getenv("DATABASE_URL"))

	r := gin.Default()

	// Authentication routes
	authRouter := r.Group("/api/auth")
	{
		authRouter.POST("/register", auth.Register)
		authRouter.POST("/login", auth.Login)
	}

	// Websocket endpoint for game events
	r.GET("/ws", websocket.HandleConnections)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

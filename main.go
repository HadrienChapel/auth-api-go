package main

import (
    "auth-api/controllers"
    "auth-api/middlewares"
    "auth-api/utils"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    utils.ConnectDatabase()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    protected := r.Group("/protected")
    protected.Use(middlewares.AuthMiddleware())
    protected.GET("/dashboard", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Welcome to the dashboard!"})
    })

    r.Run(":8080")
}

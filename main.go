package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Create a default gin router
    router := gin.Default()

    // Define a route for the root path
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Quick and Dirty HTMX+Go",
        })
    })

    // Start the server on port 9001
    router.Run(":9001")
}

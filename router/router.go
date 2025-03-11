package router

import "github.com/gin-gonic/gin"

func Initialize() {
    // Initialize router
    router := gin.Default()

    // Initialize routes
    initializeRoutes(router)

    // listen on door 8080
    router.Run(":8080")
}

package router

import "github.com/gin-gonic/gin"

func Initialize() {
    // inicialização do router
    router := gin.Default()

    // definição de uma rota
    router.GET("/ping", func(context *gin.Context) {
        context.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // executando o router na porta 8080
    router.Run(":8080")
}

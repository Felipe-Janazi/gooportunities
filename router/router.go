package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	// r agora representa o uso do gin
	r := gin.Default()

	// Definindo uma rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rodando a api
	r.Run(":8080") // Por padrão port 8080
}

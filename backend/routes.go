package main

import "github.com/gin-gonic/gin"

func registerRoutes(r *gin.Engine) {
	r.POST("/ask", askAI)
}

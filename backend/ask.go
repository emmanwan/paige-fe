package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func askAI(c *gin.Context) {
	var req struct {
		Question string `json:"question"`
	}
	c.BindJSON(&req)

	// Retrieve chunks (mock for now)
	chunks := []string{"Onboarding process is explained in HR.pdf"}
	relevant := retrieveRelevantChunks(chunks, req.Question)

	prompt := "You are an enterprise assistant. Answer ONLY from these chunks:\n" +
		strings.Join(relevant, "\n") + "\nUser question: " + req.Question

	// FIX: use AskGemini instead of callGemini
	answer, sources := AskGemini(prompt, relevant)

	c.JSON(http.StatusOK, gin.H{
		"answer":  answer,
		"sources": sources,
	})
}

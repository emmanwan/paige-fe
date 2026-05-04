package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AskGemini(question string, docs []string) (string, []string) {
	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	fmt.Printf("Using Gemini API Key: %s\n", apiKey)
	if apiKey == "" {
		log.Println("Missing GEMINI_API_KEY environment variable")
		return "", nil
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("Gemini client error: %v", err)
		return "", nil
	}
	defer client.Close()

	// 🔎 Iterate over models using Next()
	it := client.ListModels(ctx)
	var modelName string
	for {
		m, err := it.Next()
		if err != nil {
			break // iterator exhausted
		}
		log.Printf("Available model: %s", m.Name)
		if strings.Contains(m.Name, "gemini") && !strings.Contains(m.Name, "vision") {
			modelName = m.Name
			break
		}
	}

	if modelName == "" {
		log.Println("No suitable Gemini text model found")
		return "", nil
	}

	log.Printf("Using model: %s", modelName)
	model := client.GenerativeModel(modelName)

	// Build context from docs
	contextText := strings.Join(docs, "\n\n")
	prompt := "Answer the following question based only on the provided documents:\n\n" +
		contextText + "\n\nQuestion: " + question

	// Generate response
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Gemini request error: %v", err)
		return "", nil
	}

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if textPart, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(textPart), []string{"Drive Docs"}
		}
	}

	return "", nil
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/api/drive/v3"
)

type AskResponse struct {
	Answer  string   `json:"answer"`
	Sources []string `json:"sources,omitempty"`
}

var driveService *drive.Service

func main() {
	// Initialize Drive service
	driveService = initDriveService()

	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "PAIGE backend is running!")
	}).Methods("GET")

	// Ask endpoint
	r.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Question string `json:"question"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// FIX: use ENV properly
		folderId := os.Getenv("DRIVE_FOLDER_ID")
		if folderId == "" {
			http.Error(w, "Missing DRIVE_FOLDER_ID", http.StatusInternalServerError)
			return
		}

		docs, err := ListAndExportDocs(driveService, folderId)
		if err != nil {
			http.Error(w, "Error reading Drive docs: "+err.Error(), http.StatusInternalServerError)
			return
		}

		answer, sources := AskGemini(req.Question, docs)
		if answer == "" {
			answer = "I couldn’t find a direct answer in the documents. Please file a concern via Google Forms."
		}

		resp := AskResponse{Answer: answer, Sources: sources}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}).Methods("POST")

	// Enable CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200",
			"https://paige-fe.vercel.app",
		},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Authorization"},
		AllowCredentials: true,
	}).Handler(r)

	// Dynamic port (Render)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

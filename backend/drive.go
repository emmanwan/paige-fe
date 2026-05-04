package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func initDriveService() *drive.Service {
	ctx := context.Background()

	creds := os.Getenv("GOOGLE_CREDENTIALS")
	fmt.Printf("Using Google Credentials: %s\n", creds)
	if creds == "" {
		log.Fatal("GOOGLE_CREDENTIALS is empty")
	}

	service, err := drive.NewService(ctx,
		option.WithCredentialsJSON([]byte(creds)),
		option.WithScopes(drive.DriveReadonlyScope),
	)
	if err != nil {
		log.Fatalf("Unable to create Drive client: %v", err)
	}

	return service
}

func ListAndExportDocs(srv *drive.Service, folderId string) ([]string, error) {
	r, err := srv.Files.List().
		Q("'" + folderId + "' in parents").
		Fields("files(id, name, mimeType)").
		Do()

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve files: %v", err)
	}

	var docs []string

	for _, f := range r.Files {

		if f.MimeType == "application/vnd.google-apps.document" {

			res, err := srv.Files.Export(f.Id, "text/plain").Download()
			if err != nil {
				log.Printf("Unable to export file %s: %v", f.Name, err)
				continue
			}

			body, err := io.ReadAll(res.Body)
			res.Body.Close()

			if err != nil {
				log.Printf("Read error for %s: %v", f.Name, err)
				continue
			}

			chunks := chunkText(string(body), 1000)
			docs = append(docs, chunks...)

			log.Printf("Exported %s with %d chunks", f.Name, len(chunks))
		}
	}

	log.Printf("Total chunks exported: %d", len(docs))
	return docs, nil
}

func GetSheetsData(sheetID string, readRange string) ([][]interface{}, error) {
	ctx := context.Background()

	creds := os.Getenv("GOOGLE_CREDENTIALS")
	if creds == "" {
		return nil, fmt.Errorf("GOOGLE_CREDENTIALS is empty")
	}

	srv, err := sheets.NewService(ctx,
		option.WithCredentialsFile("credentials.json"),
		option.WithScopes(sheets.SpreadsheetsReadonlyScope),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create Sheets client: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(sheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	return resp.Values, nil
}

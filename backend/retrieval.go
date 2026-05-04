package main

import "strings"

func retrieveRelevantChunks(chunks []string, query string) []string {
    var relevant []string
    for _, c := range chunks {
        if strings.Contains(strings.ToLower(c), strings.ToLower(query)) {
            relevant = append(relevant, c)
        }
    }
    return relevant
}

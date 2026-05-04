package main

// chunkText splits long text into smaller pieces for Gemini
func chunkText(text string, size int) []string {
    var chunks []string
    for start := 0; start < len(text); start += size {
        end := start + size
        if end > len(text) {
            end = len(text)
        }
        chunks = append(chunks, text[start:end])
    }
    return chunks
}

package knowledgebase

import (
	"fmt"
	"os"
	"path/filepath"
)

const ChunkSize = 500

type Document struct {
	Name    string
	Content string
}

type Chunk struct {
	Content   string
	SourceDoc string
}

func LoadKnowledge() error {
	docs, err := loadMDFiles() // where to use docs??
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}
	fmt.Println("Loaded documents:", len(docs))

	for _, item := range docs {
		chunks, err := SplitDocInChunks(item)
		if err != nil {
			return fmt.Errorf("failed to split document in chunks: %w", err)
		}
		fmt.Println("Document:", item.Name)
		fmt.Println("Chunks:", len(chunks))
	}
	return nil
}

func loadMDFiles() ([]Document, error) {
	dir, err := os.ReadDir("internal/knowledge")
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}

	var documents []Document

	for _, item := range dir {
		// Read Files
		file, err := os.ReadFile(filepath.Join("internal/knowledge", item.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to load a file: %w", err)
		}
		fmt.Println("File Loaded: ", item.Name())
		// Make Empty Document Struct
		document := Document{}

		// Fill its Fields
		document.Name = item.Name()
		document.Content = string(file)

		// Append to Slice
		documents = append(documents, document)
	}

	return documents, nil
}

func SplitDocInChunks(doc Document) ([]Chunk, error) {
	// Get the Length f Document ??
	lengthDoc := len(doc.Content)

	// Make Empty Chunk slice
	chunks := []Chunk{}
	for start := 0; start < lengthDoc; start += ChunkSize {
		// Get the end
		end := start + ChunkSize
		if end > lengthDoc {
			end = lengthDoc
		}
		chunk := Chunk{
			Content:   doc.Content[start:end],
			SourceDoc: doc.Name,
		}
		chunks = append(chunks, chunk)
	}
	return chunks, nil
}

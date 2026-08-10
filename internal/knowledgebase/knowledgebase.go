package knowledgebase

import (
	"fmt"
	"os"
	"path/filepath"
)

type Document struct {
	Name    string
	Content string
}

func LoadKnowledge() error {
	docs, err := loadMDFiles() // where to use docs??
	if err != nil {
		return fmt.Errorf("failed to load docs: %w", err)
	}
	fmt.Println("Loaded documents:", len(docs))
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

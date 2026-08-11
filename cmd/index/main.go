package main

import (
	"ai-linux-cmd-assistant/internal/knowledgebase"
	"ai-linux-cmd-assistant/internal/mongodb"
	"context"
	"fmt"
)

func main() {
	client, err := mongodb.ConnectAtlasDB()
	if err != nil {
		fmt.Print("an error occured: ", err)
		return
	}
	defer client.Disconnect(context.Background())

	docs, err := knowledgebase.LoadDocuments()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, doc := range docs {
		// Delete old chunks for this document
		err := mongodb.DeleteChunks(client, doc.Name)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		// Split Document
		chunks, err := knowledgebase.SplitDocInChunks(doc)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		// Insert Fresh Chunks of Document
		err = mongodb.InsertChunks(client, chunks)
		if err != nil {
			fmt.Print("Error: ", err)
			return
		}
	}
}
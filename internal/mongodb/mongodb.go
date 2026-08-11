package mongodb

import (
	"ai-linux-cmd-assistant/internal/knowledgebase"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectAtlasDB() (*mongo.Client, error) {
	// NEED TO GET MONGO ATLAS URI FROM ENV
	// READ .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	mongoAtlasURI := os.Getenv("MONGO_ATLAS_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoAtlasURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		// panic(err)
		return nil, fmt.Errorf("an error occured connecting to mongo atlas: %w", err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// panic(err)
		return nil, fmt.Errorf("an error occured pinging client: %w", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Return Client
	return client, nil
}

func InsertChunks(client *mongo.Client, chunks []knowledgebase.Chunk) error {
	db := client.Database("ai_linux_assistant")
	collection := db.Collection("chunks")
	ctx := context.Background()
	// Loop Over and Insert Each Chunk to DB
	for _, chunk := range chunks {
		_, err := collection.InsertOne(ctx, chunk)
		if err != nil {
			return fmt.Errorf("an error occured inserting chunk into db: %w", err)
		}
	}
	return nil
}

func DeleteChunks(client *mongo.Client, sourceDoc string) error {
	db := client.Database("ai_linux_assistant")
	collection := db.Collection("chunks")
	ctx := context.Background()

	// Make Filter
	filter := bson.M{
		"source_doc": sourceDoc,
	}

	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("an error occured when deleteing chunks: %w", err)
	}
	return nil
}

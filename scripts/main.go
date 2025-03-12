package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// Read a file object from a GCloud bucket and return its contents as a byte array
func ReadFile(ctx context.Context, client *storage.Client, bucketName, objectName string) ([]byte, error) {
	// Create a reader and attach it to the given file object
	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %w", err)
	}
	// At the end, close the reader connection
	defer rc.Close()

	// Read in the byte array
	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %w", err)
	}
	return data, nil
}

// Open a GCloud storage client, access a bucket, and read its contents
// object by object using the ReadFile function
func main() {
	// Set up context
	ctx := context.Background()
	// Create client
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// At the end, close the client connection
	defer client.Close()

	bucketName := "shane-test-bucket-c77ea49a"
	// Set up bucket handle
	bkt := client.Bucket(bucketName)
	// Iterate over the objects in the bucket
	it := bkt.Objects(ctx, nil)
	for {
		obj, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to list objects: %v", err)
		}
		fmt.Println("file name:     ", obj.Name)
		// Read the file given its name
		out, err := ReadFile(ctx, client, bucketName, obj.Name)
		if err != nil {
			log.Fatalf("failed to read object: %v", err)
		}
		// Print to stdout
		fmt.Println("file contents: ", string(out))
	}

}

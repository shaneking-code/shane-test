package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func readFile(ctx context.Context, client *storage.Client, bucketName, objectName string) ([]byte, error) {
	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create reader: %w", &err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %w", &err)
	}
	return data, nil
}

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	bucketName := "shane-test-bucket-c77ea49a"
	bkt := client.Bucket(bucketName)
	it := bkt.Objects(ctx, nil)
	for {
		obj, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to list objects: %v", err)
		}
		fmt.Println("filename: ", obj.Name)
		out, err := readFile(ctx, client, bucketName, obj.Name)
		if err != nil {
			log.Fatalf("failed to read object: %v", err)
		}
		fmt.Println("file contents: ", string(out))
	}

}

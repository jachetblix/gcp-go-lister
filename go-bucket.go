package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"os"
	"time"
)

func listFiles(w io.Writer, bucket string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %w", bucket, err)
		}
		fmt.Fprintln(w, attrs.Name)
	}
	return nil
}
func main() {
	bucketName := "crossplane-bucket-8jqnx"
	if err := listFiles(os.Stdout, bucketName); err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}
}

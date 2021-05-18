package p

import (
	"context"
	"encoding/base64"
	"fmt"

	"cloud.google.com/go/storage"
)

//UploadFile function
func UploadFile(FileBase64 string, NamaFile string) error {
	ctx := context.Background()
	gcsBucketName := "afuwwu-bucket"
	// decode base64 string
	dec, err := base64.StdEncoding.DecodeString(FileBase64)
	if err != nil {
		return err
	}

	// create google cloud storage client
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// create writer and write to google cloud storage
	obj := fmt.Sprintf("test-upload/%s", NamaFile)
	wc := client.Bucket(gcsBucketName).Object(obj).NewWriter(ctx)
	if _, err := wc.Write(dec); err != nil {
		return fmt.Errorf("Writer.Write: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

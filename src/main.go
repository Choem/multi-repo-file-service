package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"file_service/handlers"
)

func main() {
	ctx := context.Background()

	endpoint := "minio"
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	useSSLString := os.Getenv("MINIO_USE_SSL")

	useSSL, useSSLConvErr := strconv.ParseBool(useSSLString)
	if useSSLConvErr != nil {
		useSSL = false
	}

	// Initialize minio client object
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// Routes
	http.HandleFunc("/create-bucket", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateBucket(w, r, minioClient, ctx)
	})
	http.HandleFunc("/remove-bucket", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemoveBucket(w, r, minioClient, ctx)
	})
	http.HandleFunc("/upload-file", func(w http.ResponseWriter, r *http.Request) {
		handlers.UploadFile(w, r, minioClient, ctx)
	})

	// Start http server
	log.Fatal(http.ListenAndServe(":4000", nil))
}

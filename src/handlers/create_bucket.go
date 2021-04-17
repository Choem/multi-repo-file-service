package handlers

import (
	"context"
	"encoding/json"
	"file_service/common"
	"log"
	"net/http"

	"github.com/minio/minio-go/v7"
)

type CreateBucketRequest struct {
	BucketName string
}

func CreateBucket(w http.ResponseWriter, r *http.Request, m *minio.Client, ctx context.Context) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		data := &CreateBucketRequest{}

		decodeErr := json.NewDecoder(r.Body).Decode(&data)
		if decodeErr != nil {
			http.Error(w, decodeErr.Error(), http.StatusBadRequest)
			return
		}

		bucketName := data.BucketName
		makeBucketErr := m.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "eu-west-1"})
		if makeBucketErr != nil {
			bucketExists, bucketExistsErr := m.BucketExists(ctx, bucketName)
			if bucketExistsErr == nil && bucketExists {
				response := &common.ApiResponse{Message: "The bucket " + bucketName + " already exists", Status: http.StatusBadRequest}
				json.NewEncoder(w).Encode(response)
			} else {
				log.Fatalln(makeBucketErr)
			}
		} else {
			response := &common.ApiResponse{Message: "Succesfully created " + bucketName, Status: http.StatusCreated}
			json.NewEncoder(w).Encode(response)
		}
	} else {
		http.Error(w, "Http method not allowed", http.StatusMethodNotAllowed)
	}
}

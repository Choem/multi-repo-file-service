package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"file_service/common"

	"github.com/minio/minio-go/v7"
)

type RemoveBucketRequest struct {
	BucketName string
}

func RemoveBucket(w http.ResponseWriter, r *http.Request, m *minio.Client, ctx context.Context) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		data := &RemoveBucketRequest{}

		decodeErr := json.NewDecoder(r.Body).Decode(&data)
		if decodeErr != nil {
			http.Error(w, decodeErr.Error(), http.StatusBadRequest)
			return
		}

		bucketName := data.BucketName
		bucketExists, bucketExistsErr := m.BucketExists(ctx, bucketName)
		if bucketExistsErr == nil && bucketExists {
			removeBucketErr := m.RemoveBucket(ctx, bucketName)
			if removeBucketErr == nil {
				response := &common.ApiResponse{Message: "Succesfully removed " + bucketName, Status: http.StatusOK}
				json.NewEncoder(w).Encode(response)
			} else {
				response := &common.ApiResponse{Message: "Could not remove " + bucketName, Status: http.StatusBadRequest}
				json.NewEncoder(w).Encode(response)
			}
		} else {
			log.Fatalln(bucketExistsErr)
		}
	} else {
		http.Error(w, "Http method not allowed", http.StatusMethodNotAllowed)
	}
}

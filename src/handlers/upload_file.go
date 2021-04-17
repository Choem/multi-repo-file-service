package handlers

import (
	"context"
	"encoding/json"
	"file_service/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

type UploadFileRequest struct {
	BucketName string
}

func UploadFile(w http.ResponseWriter, r *http.Request, m *minio.Client, ctx context.Context) {
	file, handler, fileErr := r.FormFile("file")

	if fileErr != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(fileErr)
		return
	}
	defer file.Close()

	data := &UploadFileRequest{}

	decodeErr := json.NewDecoder(r.Body).Decode(&data)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	tempDir := "/tmp/"
	tempFile, tempFileErr := ioutil.TempFile(tempDir, handler.Filename)
	if tempFileErr != nil {
		http.Error(w, tempFileErr.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	formattedTime := time.Now().Format("YYY.MM.DD.hh.mm.ss")
	extension := strings.SplitAfter(handler.Filename, ".")[0]
	fileName := formattedTime + extension
	uploadInfo, uploadErr := m.FPutObject(ctx, data.BucketName, fileName, tempDir+handler.Filename, minio.PutObjectOptions{ContentType: handler.Header.Get("mimetype")})
	if uploadErr != nil {
		http.Error(w, tempFileErr.Error(), http.StatusBadRequest)
		return
	}

	response := &common.ApiResponse{Message: "Sucessfully uploaded " + fileName + " of size " + strconv.FormatInt(uploadInfo.Size, 10), Status: http.StatusOK}
	json.NewEncoder(w).Encode(response)
}

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (e *EndPointHandler) ListBuckets(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()

	// buckets := MinioListBuckets(e.minioClient, ctx)
	buckets := S3ListBuckets(e.s3Client)
	fmt.Print(buckets)
	json.NewEncoder(w).Encode(buckets)
}

func (e *EndPointHandler) RemoveBucket(w http.ResponseWriter, r *http.Request, name string) {
	ctx := context.Background()

	resp := MinioRemoveBucket(e.minioClient, ctx, name)
	json.NewEncoder(w).Encode(resp)
}

func (e *EndPointHandler) ListObjects(w http.ResponseWriter, r *http.Request, name string) {
	// Impliment ListObject
	resp := S3ListObjects(e.s3Client, name)
	json.NewEncoder(w).Encode(resp)
}

func (e *EndPointHandler) CreateBucket(w http.ResponseWriter, r *http.Request, name string) {
	// ctx := context.Background()
	// resp := MinioMakeBucket(e.minioClient, ctx, name)

	resp := S3CreateBucket(e.s3Client, name)
	json.NewEncoder(w).Encode(resp)
}

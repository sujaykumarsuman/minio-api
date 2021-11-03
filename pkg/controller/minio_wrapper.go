package controller

import (
	"context"
	"fmt"
	"log"
	"minio-api/api"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMinio(endpoint, accessKeyID, secretAccessKey string) *minio.Client {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})

	if err != nil {
		log.Fatalf("[ERROR]: %v", err)
		os.Exit(1)
	}

	return minioClient
}

func MinioMakeBucket(mc *minio.Client, ctx context.Context, bucketName string) api.ApiResp {
	err := mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		log.Printf("[ERROR(cannot create bucket)]: %v", err)
		msg := err.Error()
		return api.ApiResp{Status: "failed", Message: msg}
	}
	msg := fmt.Sprintf("Successfully created bucket: %v", bucketName)
	return api.ApiResp{Status: "success", Message: msg}
}

func MinioListBuckets(mc *minio.Client, ctx context.Context) []minio.BucketInfo {
	buckets, err := mc.ListBuckets(ctx)
	if err != nil {
		log.Fatalf("[ERROR]: %v", err)
	}
	return buckets
}

func MinioBucketExists(mc *minio.Client, ctx context.Context, bucketName string) (bool, error) {
	found, err := mc.BucketExists(ctx, bucketName)
	if err != nil {
		return found, err
	}
	return found, nil
}

func MinioRemoveBucket(mc *minio.Client, ctx context.Context, bucketName string) api.ApiResp {
	err := mc.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		msg := err.Error()
		return api.ApiResp{Status: "failed", Message: msg}
	}
	msg := fmt.Sprintf("bucket %v successfully deleted", bucketName)
	return api.ApiResp{Status: "success", Message: msg}
}

func MinioListObjects(mc *minio.Client, ctx context.Context, cancel context.CancelFunc, bucketName string) []api.Object {
	var objectList []api.Object

	return objectList
}

package main

import (
	"minio-api/pkg/server"
	"os"
)

func main() {
	// MinIO server config
	endpoint := "minio:9000"
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")

	// Start minio-api server
	server.StartServer(endpoint, accessKeyID, secretAccessKey)
}

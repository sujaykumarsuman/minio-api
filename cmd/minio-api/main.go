package main

import "minio-api/pkg/server"

func main() {
	// MinIO server config
	endpoint := "localhost:9000"
	accessKeyID := "minio_test_db"
	secretAccessKey := "sujaykumarsuman"

	// Start minio-api server
	server.StartServer(endpoint, accessKeyID, secretAccessKey)
}

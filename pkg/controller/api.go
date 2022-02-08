package controller

import (
	"minio-api/api"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/chi/v5"
	"github.com/minio/minio-go/v7"
)

type EndPointHandler struct {
	minioClient *minio.Client
	s3Client    *s3.S3
}

func GetMinioClient(endpoint, accessKeyID, secretAccessKey string) *EndPointHandler {
	return &EndPointHandler{
		minioClient: ConnectMinio(endpoint, accessKeyID, secretAccessKey),
		s3Client:    GetNewS3Client(endpoint, accessKeyID, secretAccessKey),
	}
}

func (e *EndPointHandler) BindRequest(mux *chi.Mux) {
	mux.Route("/", func(r chi.Router) {
		api.HandlerFromMux(e, r)
	})
}

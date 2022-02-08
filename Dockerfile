FROM golang:1.17.0-alpine3.14

WORKDIR /go/src/github.com/sujaykumarsuman/minio-api

COPY . ./

RUN  go build -o /minio-api cmd/minio-api/main.go

EXPOSE 8080

CMD [ "/minio-api" ]

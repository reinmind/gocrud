package miniocli

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func New() *minio.Client {
	if client != nil {
		return client
	}
	ep := "localhost:9000"
	accessKey := "minio"
	accessSecret := "minio123"
	useSSL := false

	c, err := minio.New(ep, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: useSSL,
	})
	client = c
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("minio connected: %v", client)
	}
	return client
}

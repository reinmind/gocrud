package miniocli

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var client *minio.Client

func init() {
	ep := "localhost:9000"
	accessKey := "minio"
	accessSecret := "minio123"

	client, err := minio.New(ep, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("minio connected: %v", client)
	}
}
func Conn() *minio.Client {
	return client
}

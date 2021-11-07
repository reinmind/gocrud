package fs

import (
	_ "gocrud/config"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var client *minio.Client

func init() {
	m := viper.GetStringMapString("minio")
	var err error
	client, err = minio.New(m["endpoint"], &minio.Options{
		Creds:  credentials.NewStaticV4(m["access_key"], m["access_secret"], ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalf("%v\n", err)
		panic(err)
	}

	//log.Printf("minio configured! endpoint: %v access_key: %v access_secret: %v  \n", m["endpoint"], m["access_key"], m["access_secret"])

}

func GetClient() *minio.Client {
	return client
}

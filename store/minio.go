package store

import "os"
import "mime/multipart"
import "github.com/minio/minio-go"

type MinioFile struct {
	uri string
	object multipart.File

	client *minio.Client
}

func NewMinio(bucketName string, file multipart.File) (*MinioFile, error) {
	var mfile MinioFile
	var err error
	endpoint, access_key, secret_key  := 
		"play.minio.io:9000",
		"Q3AM3UQ867SPQQA43P2F",
		"zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	
	mfile.object = file
	mfile.client, err = minio.New(endpoint, access_key, secret_key, true)
	if err == nil {
		mfile.client, err = bucket(mfile.client, bucketName)
	}
	return &mfile, err
}

func bucket(client *minio.Client, bucketName string) (*minio.Client, error) {
	status, err := client.BucketExists(bucketName)
	if err == nil && !status {
		err = client.MakeBucket(bucketName, os.Getenv("region"))
	}
	return client, err
}

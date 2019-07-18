package store

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/minio/minio-go"
)

// MinioFile file object to be stored
type MinioFile struct {
	objectName string
	object     multipart.File
	client     *minio.Client
}

/**
//TODO 	Hard coding of credentials is bad way of practicing,
		prefer other way with respect to your architecture

Replace these credentials with your S3 compatible storage server
*/
const (
	endpoint  = "play.minio.io:9000"
	accessKey = "Q3AM3UQ867SPQQA43P2F"
	secretKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
)

// NewMinio to create new storage instance
func NewMinio(objectName string, file multipart.File) (*MinioFile, error) {
	var (
		mfile MinioFile
		err   error
	)

	// for file upload
	if file != nil {
		mfile.object = file
		defer file.Close()
	}
	mfile.objectName = objectName
	mfile.client, err = minio.New(endpoint, accessKey, secretKey, true)
	return &mfile, err
}

func bucket(client *minio.Client, bucketName string) (bool, error) {
	status, err := client.BucketExists(bucketName)
	if err == nil && !status {
		// instead of storing in environment variable, if possible
		// use other way depending on your architecture
		if client.MakeBucket(bucketName, os.Getenv("region")) != nil {
			log.Fatal("Error creating bucket", err)
		}
	}
	return status, err
}

// Upload to upload file to object storage server
func (minioFile *MinioFile) Upload(bucketName string) string {
	if status, err := bucket(minioFile.client, bucketName); err != nil && !status {
		log.Println("Error ", err)
		return "Internal Error of uploading your file"
	}

	_, err := minioFile.client.PutObject(bucketName, minioFile.objectName, minioFile.object, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		log.Println("Error in uploading ", err)
		return "Failed to upload your file"
	}
	return "File Uploaded Successfully!"
}

// Download file from object storage server
func (minioFile *MinioFile) Download(bucketName string) (*minio.Object, error) {
	return minioFile.client.GetObject(bucketName, minioFile.objectName, minio.GetObjectOptions{})
}

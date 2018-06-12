package store

import (
	"io"
	"log"
	"mime/multipart"
	"github.com/minio/minio-go"
)

type MinioFile struct {
	objectName string
	object multipart.File
	client *minio.Client
}

/**
//TODO 	Hard coding of credentials is bad way of practicing,
		prefer other way with respect to your architecture

Replace these credentials your S3 compatible storage server
*/
var (
	endpoint = "play.minio.io:9000"
	access_key = "Q3AM3UQ867SPQQA43P2F"
	secret_key = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
)

/**
* Exported function which creates storage connection
* and checks for bucket existance
*/
func NewMinio(objectName string, file multipart.File) (*MinioFile, error) {
	var (
		mfile MinioFile
		err error
	)

	// for file upload
	if file != nil {
		mfile.object = file
		defer file.Close()
	}
	mfile.objectName = objectName
	mfile.client, err = minio.New(endpoint, access_key, secret_key, true)
	return &mfile, err
}

func bucket(client *minio.Client, bucketName string) (error, bool) {
	status, err := client.BucketExists(bucketName)
	if err == nil && !status {
		// instead of storing in environment variable
		// use other way depending on your architecture
		if client.MakeBucket(bucketName, "us-east-1") != nil {
			log.Fatal("Error creating bucket", err)
		}
	}
	return err, status
}


/**
* Uploading file to storage server
* @public
*/
func (minioFile *MinioFile) Upload(bucketName string) (string) {
	if err, status := bucket(minioFile.client, bucketName); err != nil {
		if !status && err != nil {
			log.Println("Error ", err)
			return "Internal Error of uploading your file"
		}

		if _, err = minioFile.client.PutObject(bucketName, minioFile.objectName, minioFile.object, -1, minio.PutObjectOptions{ContentType: "image/jpeg"});
			 err != nil {
			log.Println("Error in uploading ", err)
			return "Failed to upload your file"
		}
	}
	return "File Uploaded Successfully!"
}

/**
* Downloading file from storage server
* @public
*/
func (minioFile *MinioFile) Download(bucketName string) ([]byte, error) {
	fileObject, err := minioFile.client.GetObject(bucketName, minioFile.objectName, minio.GetObjectOptions{})
	var fileData = make([]byte, 35000) // starting with 10kb
	var offset = int64(0)
	for err == nil {
		n, err := fileObject.ReadAt(fileData, offset)
		offset += int64(n)
		if err == io.EOF || err != nil {
			break
		}
	}
	defer fileObject.Close()
	
	return fileData, err
}

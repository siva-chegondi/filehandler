package file

import "mime/multipart"
/**
* Using Minio Go Sdk
* for Object Operations
*/
type File struct {
	objectUri string
	object multipart.File
}

/**
* Construct File with data
*/
func New(file multipart.File) (*File, int) {
	return new(File), 0
}

/**
* Uploading file to S3/Minio/Ceph
* Block Storages
*/
func (file *File) Upload() string {
	return "Uploading baby..."
}

/**
* Downloading file from S3/Minio/Ceph
* Block Storages
*/
func (file *File) Download() (*File, int) {
	return file, 0
}

package upload

import (
	"context"
	"file_flow/global"
	"github.com/minio/minio-go/v7"
	"github.com/qingstor/go-mime"
	"github.com/spf13/viper"
	"mime/multipart"
)

func PutObject(objectName string, file *multipart.FileHeader) (minio.UploadInfo, error) {
	fileReader, err := file.Open()
	if err != nil {
		return minio.UploadInfo{}, err
	}
	size := file.Size
	mimeType := mime.DetectFilePath(file.Filename)

	bucketName := viper.GetString("minio.bucketName")
	info, err := global.Minio.PutObject(context.Background(), bucketName, objectName, fileReader, size, minio.PutObjectOptions{
		ContentType: mimeType,
	})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

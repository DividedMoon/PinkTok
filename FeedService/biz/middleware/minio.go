package middleware

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7/pkg/credentials"

	"feed_service/biz/constants"
)

// minio是一个开源的对象存储服务器，
//用于存储和管理大量结构化和非结构化数据，例如照片、视频、文档等。
//它允许开发人员在本地部署或云环境中构建自己的私有云存储解决方案。
var (
	Client *minio.Client
	err    error
)

// MakeBucket 创建特定名称的桶
func MakeBucket(ctx context.Context, bucketName string) {
	exists, err := Client.BucketExists(ctx, bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !exists {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Successfully created mybucket %v\n", bucketName)
	}
}

// PutToBucket 使用 *multipart.FileHeader 将文件放到桶内
func PutToBucket(ctx context.Context, bucketName string, file *multipart.FileHeader) (info minio.UploadInfo, err error) {
	fileObj, _ := file.Open()
	info, err = Client.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	fileObj.Close()
	return info, err
}

// GetObjURL 获取minio中文件的原始url
func GetObjURL(ctx context.Context, bucketName, filename string) (u *url.URL, err error) {
	exp := time.Hour * 24
	reqParams := make(url.Values)
	u, err = Client.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return u, err
}

// PutToBucketByBuf 使用*bytes.Buffer将文件放入桶内
func PutToBucketByBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = Client.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return info, err
}

// PutToBucketByFilePath 使用filepath将文件放入桶内
func PutToBucketByFilePath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = Client.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return info, err
}

func Init() {
	ctx := context.Background()
	Client, err = minio.New(constants.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyID, constants.MinioSecretAccessKey, ""),
		Secure: constants.MiniouseSSL,
	})
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}

	log.Printf("%#v\n", Client)

	MakeBucket(ctx, constants.MinioVideoBucketName)
	MakeBucket(ctx, constants.MinioImgBucketName)
}

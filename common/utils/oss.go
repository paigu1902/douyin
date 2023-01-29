package utils

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"paigu1902/douyin/constants"
	"path/filepath"
)

func Upload(data []byte, filename string) (string, error) {
	bucketName := constants.OSSBucketName
	endpoint := constants.OSSEndpoint
	accessKeyId := constants.OSSAccessKeyId
	accessKeySecret := constants.OSSAccessKeySecret
	domain := "videos"

	client, err := oss.New("https://"+endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(filepath.Join(domain, filename), bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	//https://paigu-douyin-videos.oss-cn-nanjing.aliyuncs.com/videos/0.mp4
	return fmt.Sprintf("https://%s.%s/%s/%s", bucketName, endpoint, domain, filename), nil
}

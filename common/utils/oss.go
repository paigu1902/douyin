package utils

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"path/filepath"
)

func Upload(data []byte, filename string) (string, error) {
	bucketName := "paigu-douyin-videos"
	endpoint := "https://oss-cn-nanjing.aliyuncs.com"
	accessKeyId := "LTAI5tKz8Jg67pok7LjZ3rus"
	accessKeySecret := "BUpR46wYSyON68Qpy7aGwEvCSCvzxf"
	domain := "videos/"

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
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
	return filepath.Join(endpoint, domain, filename), nil
}

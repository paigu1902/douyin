package utils

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"paigu1902/douyin/common/config"
	"path/filepath"
)

// Upload 上传文件到阿里OSS
func Upload(data interface{}, filename string, domain string) (string, error) {
	bucketName := config.C.OSS.BucketName
	endpoint := config.C.OSS.Endpoint
	accessKeyId := config.C.OSS.AccessKeyId
	accessKeySecret := config.C.OSS.AccessKeySecret

	client, err := oss.New("https://"+endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Println(err)
		return "", err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	if val, ok := data.([]byte); ok {
		err = bucket.PutObject(filepath.Join(domain, filename), bytes.NewReader(val))
	} else if val, ok := data.(string); ok {
		err = bucket.PutObjectFromFile(filepath.Join(domain, filename), val)
	} else {
		return "", fmt.Errorf("unsupport upload type")
	}

	if err != nil {
		return "", err
	}
	//https://paigu-douyin-videos.oss-cn-nanjing.aliyuncs.com/videos/0.mp4
	return fmt.Sprintf("https://%s.%s/%s/%s", bucketName, endpoint, domain, filename), nil
}

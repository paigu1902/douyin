package main

import (
	"log"
	"time"
)

func main() {
	//c, err := userrelation.NewClient("userRelation", client.WithResolver(resolver.NewNacosResolver(nacos.Cli)))
	//if err != nil {
	//	panic(err)
	//}
	//ctx := context.Background()
	//action, err := c.FollowAction(ctx, &userRelationPb.FollowActionReq{FromId: 10, ToId: 13, Type: "0"})
	//if err != nil {
	//	return
	//}
	//fmt.Println(action.GetStatusMsg())
	//fmt.Println(action.GetStatusCode())

	// 上传视频
	//url := "http://172.23.31.167:3002/v3/action"
	//method := "POST"
	//
	//payload := &bytes.Buffer{}
	//writer := multipart.NewWriter(payload)
	//file, errFile1 := os.Open("public/bear.mp4")
	//defer file.Close()
	//base_path := filepath.Base("public/bear.mp4")
	//fmt.Println(base_path)
	//part1, errFile1 := writer.CreateFormFile("data", base_path)
	//_, errFile1 = io.Copy(part1, file)
	////fmt.Println(payload.String())
	//if errFile1 != nil {
	//	fmt.Println(errFile1)
	//	return
	//}
	//_ = writer.WriteField("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IndxeCIsImV4cCI6MTY3NjQ3NTE2Nn0.-wSI6cQH08Pr9D--lb5SlWhxDTAPb4KaEksmNu8YWCs")
	//_ = writer.WriteField("title", "carTest")
	//err := writer.Close()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	//
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	timestamp := int64(1676564215275)
	//1676564573
	//1676564215275
	//1676564629048
	//1676564605456158

	latestTime := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	log.Printf("timestamp:%d,last_time:%s, now:%d\n", timestamp, latestTime, now)
	log.Println(time.Now().UnixMilli())
}

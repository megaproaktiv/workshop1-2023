package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client
var BUCKET=os.Getenv("BUCKET")

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = s3.NewFromConfig(cfg)

}

func main() {
	path := "readme.md"
	file, err := os.Open(path)
	if err != nil {
		log.Println("Failed opening file", path, err)
	}
	defer file.Close()
	for i := 1; i < 100; i++ {
		key := fmt.Sprintf("test-go-%d.md",i)

		_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: &BUCKET,
			Key:    &key,
			Body:   file,
		})
		
		if err != nil {
			log.Println("Error:", err.Error())
			return
		}
		log.Println("File: ", key)

	}
}

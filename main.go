package main

import (
	"fmt"
	"log"

	"upload_oss/config"
)

func main() {
	config, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Println(config.Oss)
	fmt.Println(config.Oss.Endpoint)
	fmt.Println(config.Oss.AccessKeyID)
	fmt.Println(config.Oss.AccessKeySecret)
	fmt.Println(config.Oss.BucketName)
}

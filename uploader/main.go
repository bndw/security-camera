package main

import (
	"log"
	"os"
)

func main() {
	// Parse command line args
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("USAGE -- %s <image_path>\n", args[0])
	}
	filepath := args[1]

	// Load the config
	cfg := NewConfigFromEnv()
	if err := cfg.Validate(); err != nil {
		log.Println(err)
		return
	}

	// Upload the image
	fileUrl, err := UploadToS3(cfg.S3BucketRegion, cfg.S3BucketName, filepath)
	if err != nil {
		log.Printf("Failed to upload image %s to S3 with err: %s\n", filepath, err.Error())
		return
	}

	// Notify Slack
	if err := NotifySlack(cfg.SlackWebhookURL, fileUrl); err != nil {
		log.Printf("Failed to notify slack of new image %s with err: %s\n", fileUrl, err.Error())
		return
	}
}

package main

import (
	"errors"
	"os"
)

const (
	EnvSlackWebhookURL    = "SLACK_WEBHOOK_URL"
	EnvS3BucketName       = "S3_BUCKET_NAME"
	EnvS3BucketRegion     = "S3_BUCKET_REGION"
	EnvAWSAccessKeyID     = "AWS_ACCESS_KEY_ID"
	EnvAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

type Config struct {
	// Slack config
	SlackWebhookURL string
	// AWS config
	S3BucketName       string
	S3BucketRegion     string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}

func (c *Config) Validate() error {
	if c.SlackWebhookURL == "" {
		return errors.New("Must set env var " + EnvSlackWebhookURL)
	}
	if c.S3BucketName == "" {
		return errors.New("Must set env var " + EnvS3BucketName)
	}
	if c.S3BucketRegion == "" {
		return errors.New("Must set env var " + EnvS3BucketRegion)
	}
	if c.AWSAccessKeyID == "" {
		return errors.New("Must set env var " + EnvAWSAccessKeyID)
	}
	if c.AWSSecretAccessKey == "" {
		return errors.New("Must set env var " + EnvAWSSecretAccessKey)
	}

	return nil
}

func NewConfigFromEnv() *Config {
	return &Config{
		SlackWebhookURL:    os.Getenv(EnvSlackWebhookURL),
		S3BucketName:       os.Getenv(EnvS3BucketName),
		S3BucketRegion:     os.Getenv(EnvS3BucketRegion),
		AWSAccessKeyID:     os.Getenv(EnvAWSAccessKeyID),
		AWSSecretAccessKey: os.Getenv(EnvAWSSecretAccessKey),
	}
}

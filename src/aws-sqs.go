package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	// AWS SETTINGS
	accessKey = ""
	secretKey = ""
	region    = "us-west-2"
	sqsURL    = "https://sqs.us-west-2.amazonaws.com/917204115845/jor-el-wordpress-ingestion-posts"
)

/**
* Save the structured data into SQS
***/
func saveToSQS(payload string) {
	// Send the payload to the SQS.
	fmt.Println("Saving posts....")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	sqsClient := sqs.New(sess)
	qURL := sqsURL

	results, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &qURL,
		MessageBody: aws.String(payload),
	})

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(results)
}

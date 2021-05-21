package sqs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Queue struct {
	Region string
	Name string
}

func getQueueProps(q Queue) (*sqs.SQS, string) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(q.Region)},
	)
	svc := sqs.New(sess)

	url, _ := svc.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: &q.Name})
    
	return svc, *url.QueueUrl
}

func ListQueues(region string) {
	svc, _ := getQueueProps(Queue{Region: region})
	result, err := svc.ListQueues(nil)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success")

	for i, urls := range result.QueueUrls {
		if urls == nil {
			continue
		}
		fmt.Printf("%d: %s\n", i, *urls)
	}
}

package main

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	// dynamodb stuffs
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Item: Create struct to hold info about new item
type Item struct {
	Timeslice string
	SortKey   string
}

func handler(ctx context.Context, s3Event events.S3Event) (string, error) {
	/*

		// For grabbing s3 metadata when required

			svc := s3.New(session.New())

				input := &s3.HeadObjectInput{
					Bucket: aws.String(s3Event.Records[0].S3.Bucket.Name),
					Key:    aws.String(s3Event.Records[0].S3.Object.Key),
				}

				// Attempt to read the head object
				result, err := svc.HeadObject(input)
				if err != nil {
					if aerr, ok := err.(awserr.Error); ok {
						switch aerr.Code() {
						default:
							return aerr.Error(), err
						}
					} else {
						// Print the error, cast err to awserr.Error to get the Code and
						// Message from an error.
						return err.Error(), err
					}
				}*/

	// Write item to DynamoDB
	// --------------------------------------------------------------------------
	dynamosvc := dynamodb.New(session.New())
	item := Item{
		//partition: s3Event.Records[0].EventTime.UTC().Format("2006-01-02"),
		//sort:      s3Event.Records[0].EventTime.UTC().Format("15:04:05.999") + " | " + s3Event.Records[0].S3.Object.Key,
		Timeslice: "2019",
		SortKey:   "Avengers: Endgame",
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err.Error(), err
	}

	tableName := "s3data"

	dynamoinput := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamosvc.PutItem(dynamoinput)
	if err != nil {
		return err.Error(), err
	}
	// --------------------------------------------------------------------------

	return "Successfully added item to s3data table", nil

	//fmt.Println(result)
	//fmt.Println(result.Metadata)
	//return string(*result.Metadata["Date"]), nil
	//	fmt.Println(*result.Metadata["Date"])
}

func main() {
	lambda.Start(handler)
}

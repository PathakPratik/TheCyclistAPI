package AwsDynamoDb

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamoClient *dynamodb.DynamoDB

func InitDynamoDb() {
	// Create AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	_, err := sess.Config.Credentials.Get()

	if err != nil {
		log.Fatal("Aws Error:", err)
	}

	// Create DynamoDB client
	dynamoClient = dynamodb.New(sess)
}

type TrackingData struct {
	UserId    json.Number `valid:"required"`
	RecordId  json.Number `valid:"required"`
	TripId    json.Number `valid:"required"`
	Latitude  json.Number `valid:"required"`
	Longitude json.Number `valid:"required"`
	Timestamp json.Number `valid:"required"`
}

const TableName = "UserTrips"
const IndexName = "UserId"

func AddItem(item TrackingData) bool {
	av, _ := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err := dynamoClient.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem: ", err.Error())
	}

	return true
}

func GetItem(Uid string) (bool, []TrackingData) {

	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(TableName),
		IndexName: aws.String("UserId-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			IndexName: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(Uid),
					},
				},
			},
		},
	}

	var res, err = dynamoClient.Query(queryInput)

	trackingData := []TrackingData{}

	if err != nil {
		fmt.Println("Got error calling Query: ", err.Error())
		return false, trackingData
	} else {
		dynamodbattribute.UnmarshalListOfMaps(res.Items, &trackingData)
		return true, trackingData
	}
}

type Events struct {
	Eventname    string
	EventAddress string
	EventCity    string
	EventDate    string
	EventURL     string
}

func AddEvent(item Events) bool {
	av, _ := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Events"),
	}

	_, err := dynamoClient.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem: ", err.Error())
	}

	return true
}

func GetEvents() (bool, []Events) {

	var queryInput = &dynamodb.ScanInput{TableName: aws.String("Events")}

	var res, err = dynamoClient.Scan(queryInput)

	events := []Events{}

	if err != nil {
		fmt.Println("Got error calling Query: ", err.Error())
		return false, events
	} else {
		dynamodbattribute.UnmarshalListOfMaps(res.Items, &events)
		return true, events
	}
}

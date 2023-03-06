package repository

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type PetRepository struct {
	dynamoDB  *dynamodb.DynamoDB
	tableName string
}

// PetEntity defines model for Pet in the database.
type PetEntity struct {
	ID   int64   `dynamodbav:"ID"`
	Name string  `dynamodbav:"Name"`
	Tag  *string `dynamodbav:"Tag,omitempty"`
}

func NewPetRepository() *PetRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return &PetRepository{
		dynamoDB:  dynamodb.New(sess),
		tableName: "Pet",
	}
}

func (r *PetRepository) Add(pet PetEntity) error {
	item, err := dynamodbattribute.MarshalMap(pet)
	if err != nil {
		log.Printf("Error marshalling pet entity: %v", err)
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:                item,
		TableName:           aws.String(r.tableName),
		ConditionExpression: aws.String("attribute_not_exists(ID)"),
	}
	_, err = r.dynamoDB.PutItem(input)
	if err != nil {
		log.Printf("Error adding pet entity to DynamoDB: %v", err)
		return err
	}
	return nil
}

func (r *PetRepository) GetPet(id string) (*PetEntity, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(r.tableName),
	}
	result, err := r.dynamoDB.GetItem(input)
	if err != nil {
		log.Printf("Error getting pet entity from DynamoDB: %v", err)
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.New("pet not found")
	}
	var pet PetEntity
	err = dynamodbattribute.UnmarshalMap(result.Item, &pet)
	if err != nil {
		log.Printf("Error unmarshalling pet entity: %v", err)
		return nil, err
	}
	return &pet, nil
}

func (r *PetRepository) GetPets(limit uint8) ([]PetEntity, error) {
	var pets []PetEntity
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}
	result, err := r.dynamoDB.Scan(input)
	if err != nil {
		log.Printf("Error scanning pet entities from DynamoDB: %v", err)
		return nil, err
	}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &pets)
	if err != nil {
		log.Printf("Error unmarshalling list of pet entities: %v", err)
		return nil, err
	}
	if limit == 0 || limit > uint8(len(pets)) {
		limit = uint8(len(pets))
	}
	return pets[:limit], nil
}

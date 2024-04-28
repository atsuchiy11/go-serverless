package adapter

import "go_serverless/infra"

type ResourceTableOperator struct {
	TableOperator
}

type ResourceSchema struct {
	PK string `dynamo:"PK,hash"`
	SK string `dynamo:"SK,range"`
}

func NewResourceTableOperator(client *infra.DynamoClient, tableName string) *ResourceTableOperator {
	return &ResourceTableOperator{TableOperator: *NewTableOperator(client, tableName)}
}

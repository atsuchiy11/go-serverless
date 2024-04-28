package adapter

import "go_serverless/infra"

type TableOperator struct {
	Client    *infra.DynamoClient
	TableName string
}

func NewTableOperator(client *infra.DynamoClient, tableName string) *TableOperator {
	return &TableOperator{Client: client, TableName: tableName}
}

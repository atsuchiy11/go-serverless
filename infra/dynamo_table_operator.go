package infra

import (
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
	"go_serverless/domain"
)

type TableOperator struct {
	Table     *dynamo.Table
	TableName string
}

func NewTableOperator(table *dynamo.Table, tableName string) *TableOperator {
	return &TableOperator{Table: table, TableName: tableName}
}

func (o *TableOperator) CreateUser(newUser *domain.UserModel) (*domain.UserModel, error) {
	err := o.Table.Put(newUser).Run()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return newUser, nil
}

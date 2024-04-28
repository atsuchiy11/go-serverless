package registry

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
	"go_serverless/infra"
	"go_serverless/interactor"
	"go_serverless/usecase"
)

var FactorySingleton *Factory

type Factory struct {
	Envs     *Envs
	dynamoDB *dynamo.DB
}

func GetFactory() *Factory {
	if FactorySingleton == nil {
		FactorySingleton = &Factory{}
	}
	return FactorySingleton
}

// 構造体のフィールドでシングルトンを実現するパターン

func (f *Factory) NewDynamoDBClient() (*dynamo.DB, error) {
	if f.dynamoDB != nil {
		return f.dynamoDB, nil
	}
	config := &aws.Config{
		Region:   aws.String(f.Envs.AWSRegion()),
		Endpoint: aws.String(f.Envs.DynamoLocalEndpoint()),
	}

	db, err := infra.NewDynamoClient(config).Connect()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	f.dynamoDB = db
	return f.dynamoDB, nil
}

func (f *Factory) BuildTableOperator(tableName string) *infra.TableOperator {
	table := f.dynamoDB.Table(tableName)
	return infra.NewTableOperator(&table, tableName)

}

func (f *Factory) BuildCreateUser() usecase.ICreateUser {
	return interactor.NewCreateUser(f.BuildTableOperator(f.Envs.DynamoTableName()))
}

func (f *Factory) BuildCreateHelloMessage() usecase.ICreateHelloMessage {
	return interactor.NewCreateHelloMessage()
}

// コンテナでシングルトンを実現するパターン

//func (f *Factory) container(key string, builder func() interface{}) interface{} {
//	if f.cache == nil {
//		f.cache = map[string]interface{}{}
//	}
//	if f.cache[key] == nil {
//		f.cache[key] = builder()
//	}
//	return f.cache[key]
//}

//func (f *Factory) BuildDynamoClient() *infra.DynamoClient {
//	return f.container("DynamoClient", func() interface{} {
//		config := &aws.Config{
//			Region:   aws.String(f.Envs.AWSRegion()),
//			Endpoint: aws.String(f.Envs.DynamoLocalEndpoint()),
//		}
//		return infra.NewDynamoClient(config)
//	}).(*infra.DynamoClient)
//}
//
//func (f *Factory) BuildResourceTableOperator() *adapter.ResourceTableOperator {
//	return f.container("ResourceTableOperator", func() interface{} {
//		return adapter.NewResourceTableOperator(f.BuildDynamoClient(), f.Envs.DynamoTableName())
//	}).(*adapter.ResourceTableOperator)
//}

//func (f *Factory) BuildUserOperator() domain.UserRepository {
//	return f.container("UserOperator", func() interface{} {
//		return &adapter.UserOperator{Client: f.BuildResourceTableOperator()}
//	}).(domain.UserRepository)
//}

//func (f *Factory) BuildCreateUser() usecase.ICreateUser {
//	return f.container("CreateUser", func() interface{} {
//		return interactor.NewCreateUser(f.BuildUserOperator())
//	}).(usecase.ICreateUser)
//}

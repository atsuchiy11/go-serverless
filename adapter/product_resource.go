package adapter

import "go_serverless/domain"

// ProductResource ProductModelのDynamoDB用構造体
type ProductResource struct {
	ResourceSchema
	DynamoResourceBase
	domain.ProductModel
	Mapper *DynamoModelMapper `dynamo:"-"`
}

func NewProductResource(productModel *domain.ProductModel, mapper *DynamoModelMapper) *ProductResource {
	return &ProductResource{
		ProductModel: *productModel,
		Mapper:       mapper,
	}
}

// 以下、DynamoResourceインタフェースの実装

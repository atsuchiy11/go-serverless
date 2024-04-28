package adapter

import (
	"reflect"
	"time"
)

type DynamoResource interface {
	EntityName() string
	PK() string
	SetPK()
	SK() string
	SetSK()
	SetID(id uint64)
	ID() uint64
	SetVersion(v int)
	Version() int
	CreatedAt() time.Time
	SetCreatedAt(t time.Time)
	UpdatedAt() time.Time
	SetUpdatedAt(t time.Time)
}

type DynamoModelMapper struct {
	Client    *ResourceTableOperator
	TableName string
	PKName    string
	SKName    string
}

func (d *DynamoModelMapper) GetEntityNameFromStruct(s interface{}) string {
	r := reflect.TypeOf(s)
	return r.Name()
}

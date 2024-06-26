package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

type DynamoClient struct {
	Client *dynamo.DB
	Config *aws.Config
}

func NewDynamoClient(config *aws.Config) *DynamoClient {
	return &DynamoClient{Config: config}
}

func (c *DynamoClient) Connect() (*dynamo.DB, error) {
	if c.Client == nil {
		sess, err := session.NewSession(c.Config)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		c.Client = dynamo.New(sess)
	}
	return c.Client, nil
}

package registry

import "os"

type Envs struct {
	Cache map[string]string
}

var envs *Envs

func NewEnvs() *Envs {
	return &Envs{
		Cache: make(map[string]string),
	}
}

func Env() *Envs {
	if envs == nil {
		envs = NewEnvs()
	}
	return envs
}

func (c *Envs) env(key string) string {
	return os.Getenv(key)
}

func (c *Envs) AWSRegion() string {
	return c.env("AWS_REGION")
}

func (c *Envs) DynamoLocalEndpoint() string {
	return c.env("DYNAMO_LOCAL_ENDPOINT")
}

func (c *Envs) DynamoTableName() string {
	return c.env("DYNAMO_TABLE_NAME")
}

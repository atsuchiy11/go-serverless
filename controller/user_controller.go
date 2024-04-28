package controller

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"go_serverless/registry"
	"go_serverless/usecase"
)

func PostUser(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	// usecaseのパラメータに変換
	var req usecase.CreateUserRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response500(err)
	}

	i := registry.GetFactory().BuildCreateUser()
	res, err := i.Execute(&req)
	if err != nil {
		return Response500(err)
	}
	return Response201(res.GetUserID())
}

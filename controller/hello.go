package controller

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"go_serverless/registry"
	"go_serverless/usecase"
)

// PostHello
//  1. HTTPリクエストから必要なパラメータを抽出する
//  2. (バリデーションを行う)
//  3. usecaseにパラメータを渡す
//  4. usecaseからレスポンスを受け取る
//  3. HTTPレスポンスとして返す
func PostHello(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	// usecaseのパラメータに変換
	var req usecase.CreateHelloMessageRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return Response500(err)
	}

	h := registry.GetFactory().BuildCreateHelloMessage()
	res, err := h.Execute(&req)
	if err != nil {
		return Response500(err)
	}

	return Response200(res)
}

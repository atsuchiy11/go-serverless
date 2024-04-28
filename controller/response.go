package controller

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func commonHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}
}

func Response200(body interface{}) events.APIGatewayProxyResponse {
	b, err := json.Marshal(body)
	if err != nil {
		return Response500(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    commonHeaders(),
		Body:       string(b),
	}
}

func Response201(id uint64) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    commonHeaders(),
		Body:       fmt.Sprintf(`{"message":"OK","id":%d}`, id),
	}
}
func Response500(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    commonHeaders(),
		Body:       `{"message":"サーバエラーが発生しました"}`,
	}
}

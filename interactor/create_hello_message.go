package interactor

import (
	"fmt"
	"go_serverless/usecase"
)

type CreateHelloMessage struct{}

func NewCreateHelloMessage() *CreateHelloMessage {
	return &CreateHelloMessage{}
}

func (u *CreateHelloMessage) Execute(req *usecase.CreateHelloMessageRequest) (*usecase.CreateHelloMessageResponse, error) {
	msg := fmt.Sprintf("Hello %s", req.Name)
	return &usecase.CreateHelloMessageResponse{Message: msg}, nil
}

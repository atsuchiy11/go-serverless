package interactor

import (
	"github.com/pkg/errors"
	"go_serverless/domain"
	"go_serverless/usecase"
)

type CreateUser struct {
	UserRepository domain.UserRepository
}

func NewCreateUser(repos domain.UserRepository) *CreateUser {
	return &CreateUser{UserRepository: repos}
}

func (u *CreateUser) Execute(req *usecase.CreateUserRequest) (*usecase.CreateUserResponse, error) {
	user, err := u.UserRepository.CreateUser(req.ToUserModel())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &usecase.CreateUserResponse{User: user}, nil
}

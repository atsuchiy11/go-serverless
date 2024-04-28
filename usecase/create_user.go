package usecase

import "go_serverless/domain"

type ICreateUser interface {
	Execute(req *CreateUserRequest) (*CreateUserResponse, error)
}

type CreateUserRequest struct {
	Name  string
	Email string
}

func (u *CreateUserRequest) ToUserModel() *domain.UserModel {
	return domain.NewUserModel(u.Name, u.Email)
}

type CreateUserResponse struct {
	User *domain.UserModel
}

func (u *CreateUserResponse) GetUserID() uint64 {
	return u.User.ID
}

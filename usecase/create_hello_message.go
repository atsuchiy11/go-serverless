package usecase

type ICreateHelloMessage interface {
	Execute(req *CreateHelloMessageRequest) (*CreateHelloMessageResponse, error)
}

type CreateHelloMessageRequest struct {
	Name string `json:"name"`
}

type CreateHelloMessageResponse struct {
	Message string `json:"message"`
}

package handler

import (
	"context"
	"log"

	contractv1 "github.com/sadensmol/article_my_clean_architecture_go_application/api/proto/v1"
)

type AccountUsecases interface {
	Create()
}
type AccountHandler struct {
	AccountUsecases AccountUsecases
}

// here we do validate incoming requests!!!
func (h *AccountHandler) Create(ctx context.Context, request *contractv1.CreateAccountRequest) (*contractv1.CreateAccountResponse, error) {
	log.Println("create was called!!!")
	h.AccountUsecases.Create()
	return &contractv1.CreateAccountResponse{Id: "test_d"}, nil
}
func (h *AccountHandler) GetById(ctx context.Context, request *contractv1.GetAccountRequest) (*contractv1.GetAccountResponse, error) {
	log.Println("get was called!!!")
	return &contractv1.GetAccountResponse{
		Id:     "test_id",
		Name:   "test_name",
		Status: contractv1.AccountStatus_ACCOUNT_STATUS_NEW,
	}, nil
}

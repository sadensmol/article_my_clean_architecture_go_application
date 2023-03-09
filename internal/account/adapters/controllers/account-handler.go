package controllers

import (
	"context"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"

	contractv1 "github.com/sadensmol/article_my_clean_architecture_go_application/api/proto/v1"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"
)

type AccountHandler struct {
	accountUsecases AccountUsecases
}

type AccountUsecases interface {
	Create(account domain.Account) (*domain.Account, error)
	GetByID(ID int64) (*domain.Account, error)
}

func NewAccountHandler(accountUsecases AccountUsecases) *AccountHandler {
	return &AccountHandler{accountUsecases: accountUsecases}
}

func (h *AccountHandler) Create(ctx context.Context, request *contractv1.CreateAccountRequest) (*contractv1.CreateAccountResponse, error) {
	log.Println("create was called!!!")
	account, err := h.accountUsecases.Create(domain.Account{
		Name:        request.Name,
		Status:      domain.AccountStatusNew,
		AccessLevel: AccountHandlerAccessLevel(request.AccessLevel).ToDomain(),
	})

	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}

	return &contractv1.CreateAccountResponse{Id: account.ID}, nil
}
func (h *AccountHandler) GetById(ctx context.Context, request *contractv1.GetAccountRequest) (*contractv1.GetAccountResponse, error) {
	account, err := h.accountUsecases.GetByID(request.GetId())
	if err != nil {
		log.Fatalf("Error occurred %v", err)
	}

	if account == nil {
		return nil, protoregistry.NotFound
	}

	return &contractv1.GetAccountResponse{
		Id:          account.ID,
		Name:        account.Name,
		Status:      AccountStatusFromDomain(account.Status),
		AccessLevel: AccountAccessLevelFromDomain(account.AccessLevel),
		OpenedDate:  timestamppb.New(account.OpenedAt),
		ClosedDate: func() *timestamppb.Timestamp {
			if account.ClosedAt != nil {
				return timestamppb.New(*account.ClosedAt)
			} else {
				return nil
			}
		}(),
	}, nil
}

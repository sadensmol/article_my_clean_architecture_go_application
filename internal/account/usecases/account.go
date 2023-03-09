package usecases

import (
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"
)

type AccountUsecases struct {
	accountService AccountService
}

type AccountService interface {
	Create(account domain.Account) (*domain.Account, error)
	GetByID(ID int64) (*domain.Account, error)
}

func NewAccountUsecases(accountService AccountService) *AccountUsecases {
	return &AccountUsecases{accountService: accountService}
}

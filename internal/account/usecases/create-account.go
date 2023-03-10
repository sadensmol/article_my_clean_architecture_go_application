package usecases

import (
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"
)

func (a *AccountUsecases) Create(account domain.Account) (*domain.Account, error) {
	return a.accountService.Create(account)
}

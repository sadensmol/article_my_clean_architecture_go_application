package usecases

import (
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"
)

func (a *AccountUsecases) GetByID(ID int64) (*domain.Account, error) {
	return a.accountService.GetByID(ID)
}

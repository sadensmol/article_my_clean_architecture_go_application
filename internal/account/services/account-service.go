package services

import "github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"

type AccountService struct {
	accountRepository AccountRepository
}

type AccountRepository interface {
	Create(account domain.Account) (*domain.Account, error)
	GetByID(ID int64) (*domain.Account, error)
}

func NewAccountService(accountRepository AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) Create(account domain.Account) (*domain.Account, error) {
	return s.accountRepository.Create(account)
}

func (s *AccountService) GetByID(ID int64) (*domain.Account, error) {
	return s.accountRepository.GetByID(ID)
}

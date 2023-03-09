package repositories

import (
	"database/sql"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"

	"github.com/sadensmol/article_my_clean_architecture_go_application/db/gen/account/public/model"
	"github.com/sadensmol/article_my_clean_architecture_go_application/db/gen/account/public/table"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (a *AccountRepository) Create(account domain.Account) (*domain.Account, error) {

	var savedAccount model.Account
	err := table.Account.INSERT(table.Account.Name, table.Account.AccessLevel).
		VALUES(account.Name, account.AccessLevel).
		RETURNING(table.Account.AllColumns).Query(a.db, &savedAccount)

	if err != nil {
		return nil, err
	}

	return RepositoryModelAccount(savedAccount).toDomain(), nil
}

func (a *AccountRepository) GetByID(ID int64) (*domain.Account, error) {

	var savedAccounts []model.Account
	err := table.Account.SELECT(table.Account.AllColumns).
		WHERE(table.Account.ID.EQ(postgres.Int(ID))).
		Query(a.db, &savedAccounts)

	if err != nil {
		return nil, err
	}

	if len(savedAccounts) == 0 {
		return nil, nil
	}
	if len(savedAccounts) > 1 {
		return nil, fmt.Errorf("GetByID returns non unique result")
	}

	return RepositoryModelAccount(savedAccounts[0]).toDomain(), nil
}

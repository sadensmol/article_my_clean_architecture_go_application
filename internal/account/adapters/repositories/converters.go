package repositories

import (
	"github.com/sadensmol/article_my_clean_architecture_go_application/db/gen/account/public/model"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"
)

type RepositoryModelAccount model.Account

func (a RepositoryModelAccount) toDomain() *domain.Account {
	return &domain.Account{
		ID:   a.ID,
		Name: a.Name,
		Status: func() domain.AccountStatus {
			switch a.Status {
			case "open":
				return domain.AccountStatusOpen
			case "closed":
				return domain.AccountStatusClosed
			case "new":
				return domain.AccountStatusNew
			default:
				return domain.AccountStatusUnknown
			}

		}(),
		AccessLevel: func() domain.AccountAccessLevel {
			switch a.AccessLevel {
			case "full":
				return domain.AccountAccessLevelFull
			case "readonly":
				return domain.AccountAccessLevelReadOnly
			case "none":
				return domain.AccountAccessLevelNone
			default:
				return domain.AccountAccessLevelNone
			}
		}(),
		OpenedAt: a.OpenedAt,
		ClosedAt: a.ClosedAt,
	}
}

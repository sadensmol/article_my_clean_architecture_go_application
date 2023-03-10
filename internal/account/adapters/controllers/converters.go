package controllers

import (
	contractv1 "github.com/sadensmol/article_my_clean_architecture_go_application/api/proto/v1"
	"github.com/sadensmol/article_my_clean_architecture_go_application/internal/account/domain"
)

type AccountHandlerAccessLevel contractv1.AccessLevel

func (cal AccountHandlerAccessLevel) ToDomain() domain.AccountAccessLevel {
	switch contractv1.AccessLevel(cal) {
	case contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_FULL_ACCESS:
		return domain.AccountAccessLevelFull
	case contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_READ_ONLY:
		return domain.AccountAccessLevelReadOnly
	case contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_NO_ACCESS:
		return domain.AccountAccessLevelNone
	default:
		return domain.AccountAccessLevelNone
	}
}

func AccountAccessLevelFromDomain(accessLevel domain.AccountAccessLevel) contractv1.AccessLevel {
	switch accessLevel {
	case domain.AccountAccessLevelFull:
		return contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_FULL_ACCESS
	case domain.AccountAccessLevelReadOnly:
		return contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_READ_ONLY
	case domain.AccountAccessLevelNone:
		return contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_NO_ACCESS
	default:
		return contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_UNKNOWN
	}
}

func AccountStatusFromDomain(status domain.AccountStatus) contractv1.AccountStatus {
	switch status {
	case domain.AccountStatusNew:
		return contractv1.AccountStatus_ACCOUNT_STATUS_NEW
	case domain.AccountStatusOpen:
		return contractv1.AccountStatus_ACCOUNT_STATUS_OPEN
	case domain.AccountStatusClosed:
		return contractv1.AccountStatus_ACCOUNT_STATUS_CLOSED
	default:
		return contractv1.AccountStatus_ACCOUNT_STATUS_UNKNOWN
	}
}

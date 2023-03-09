package domain

import (
	"time"
)

type AccountStatus string

const (
	AccountStatusNew     AccountStatus = "new"
	AccountStatusOpen    AccountStatus = "open"
	AccountStatusClosed  AccountStatus = "closed"
	AccountStatusUnknown AccountStatus = "unknown"
)

type AccountAccessLevel string

const (
	AccountAccessLevelFull     AccountAccessLevel = "full"
	AccountAccessLevelReadOnly AccountAccessLevel = "readonly"
	AccountAccessLevelNone     AccountAccessLevel = "none"
)

type Account struct {
	ID          int64
	Name        string
	Status      AccountStatus
	AccessLevel AccountAccessLevel
	OpenedAt    time.Time
	ClosedAt    *time.Time
}

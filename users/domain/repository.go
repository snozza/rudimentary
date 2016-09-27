package users

import (
  "github.com/snozza/email-ads-data-api/domain"
)

type IUserRepositoryFactory interface {
  New(db domain.IDatabase) IUserRepository
}

type IUserRepository interface {
  FilterUsers(uuids []string, lastID string, limit int) domain.IUsers
}

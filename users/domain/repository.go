package users

import (
  "github.com/snozza/rudimentary/domain"
)

type IUserRepositoryFactory interface {
  New(db domain.IDatabase) IUserRepository
}

type IUserRepository interface {
  FilterUsers(field string, query string, lastID string, limit int, sort string) domain.IUsers
}

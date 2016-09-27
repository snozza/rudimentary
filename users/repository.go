package users

import (
  . "github.com/snozza/email-ads-data-api/users/domain"
  "github.com/snozza/email-ads-data-api/domain"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "log"
)

const UsersCollection string = "users"

func NewUserRepositoryFactory() IUserRepositoryFactory {
  return &UserRepositoryFactory{}
}

type UserRepositoryFactory struct{}

func(factory *UserRepositoryFactory) New(db domain.IDatabase) IUserRepository {
  return &UserRepository{db}
}

type UserRepository struct {
  DB domain.IDatabase
}

func (repo *UserRepository) FilterUsers(uuids []string, lastID string, limit int) domain.IUsers {

  users := Users {}

  // ensure that collection has the right text index
  // refactor building collection index
  err := repo.DB.EnsureIndex(UsersCollection, mgo.Index{
    Key: []string{
      "$text:uuid",
    },
    Background: true,
    Sparse: true,
  })
  if err != nil {
    log.Println("FilterUsers: EnsureIndex", err.Error())
  }

  q := domain.Query{
    "uuid": bson.M{"$in": uuids},
  }

  if lastID != "" && bson.IsObjectIdHex(lastID) {
    q["_id"] = domain.Query{
      "$gt": bson.ObjectIdHex(lastID),
    }
  }

  err = repo.DB.FindAll(UsersCollection, q, &users, limit)
  if err != nil {
    return &Users{}
  }
  return &users
}


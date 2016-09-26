package users

import (
  . "github.com/snozza/rudimentary/users/domain"
  "fmt"
  "github.com/snozza/rudimentary/domain"
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

func (repo *UserRepository) FilterUsers(field string, query string, lastID string, limit int, sort string) domain.IUsers {

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
  // parse sort string
  allowedSortMap := map[string]bool{
    "_id": true,
    "-_id": true,
  }
  // ensure that sort string is allowed
  // we are basically concerned about sorting on un-indexed keys
  if !allowedSortMap[sort] {
    sort = "-_id" // set it to default sort
  }

  q := domain.Query{}
  if lastID != "" && bson.IsObjectIdHex(lastID) {
    if sort == "_id" {
      q["_id"] = domain.Query{
        "$gt": bson.ObjectIdHex(lastID),
      }
    } else {
      q["_id"] = domain.Query{
        "$lt": bson.ObjectIdHex(lastID),
      }
    }
  }

  if query != "" {
    if field != "" {
      q[field] = domain.Query{
        "$regex": fmt.Sprintf("^%v.*", query),
        "$options": "i",
      }
    } else {
      // if not field is specified, we do a text search on pre-defined text index
      q["$text"] = domain.Query{
        "$search": query,
      }
    }
  }

  err = repo.DB.FindAll(UsersCollection, q, &users, limit, sort)
  if err != nil {
    return &Users{}
  }
  return &users
}


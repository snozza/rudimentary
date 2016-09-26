package domain

import (
  "gopkg.in/mgo.v2"
)

type Query map[string]interface{}
type Change mgo.Change

type IDatabase interface {
  FindAll(name string, query Query, result interface{}, limit int, sort string) error
  EnsureIndex(name string, index mgo.Index) error
}


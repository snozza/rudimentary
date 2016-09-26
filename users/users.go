package users

import (
  "gopkg.in/mgo.v2/bson"
)

type User struct {
  ID  bson.ObjectId `json:"id,omitempty" bson:"_id, omitempty"`
  Username string `json:"username,omitempty" bson:"username"`
  Email string `json:"email,omitempty" bson:"email"`
}

type Users []User

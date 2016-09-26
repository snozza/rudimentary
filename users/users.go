package users

import (
  "gopkg.in/mgo.v2/bson"
)

type User struct {
  ID  bson.ObjectId `json:"id,omitempty" bson:"_id"`
  Uuid string `json:"uuid,omitempty" bson:"uuid"`
  Fifth string `json:"05,omitempty" bson:"05"`
  Sixth string `json:"06,omitempty" bson:"06"`
  Seventh string `json:"07,omitempty" bson:"07"`
  Gender string `json:"gender,omitempty" bson:"gender"`
}

type Users []User

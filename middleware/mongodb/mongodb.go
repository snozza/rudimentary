package mongodb

import (
  "net/http"
  "github.com/snozza/rudimentary/domain"
  "gopkg.in/mgo.v2"
  "time"
)

const MongoDbKey domain.ContextKey = "rudimentary-mongodb-key"

type Options struct {
  ServerName string
  DatabaseName string
  DialTimeout time.Duration
}

func New(options *Options) *MongoDB {
  db := &MongoDB{}
  db.options = options
  return db
}

// MongoDatabase implements IDatabase
type MongoDB struct {
  currentDb *mgo.Database
  options *Options
}

func (db *MongoDB) NewSession() *MongoDBSession {
  mongoOptions := db.options

  // set default DialTimeout value

  if mongoOptions.DialTimeout <= 0 {
    mongoOptions.DialTimeout = 1 * time.Minute
  }

  session, err := mgo.DialWithTimeout(mongoOptions.ServerName, mongoOptions.DialTimeout)
  if err != nil {
    panic(err)
  }
  db.currentDb = session.DB(mongoOptions.DatabaseName)
  return &MongoDBSession{session, mongoOptions}
}

func (db *MongoDB) FindAll(name string, query domain.Query, result interface{},
  limit int, sort string) error {

  if sort == "" {
    sort = "-_id"
  }
  return db.currentDb.C(name).Find(query).Sort(sort).Limit(limit).All(result)
}

func (db *MongoDB) EnsureIndex(name string, index mgo.Index) error {
  return db.currentDb.C(name).EnsureIndex(index)
}

type MongoDBSession struct {
  *mgo.Session
  *Options
}

func (session *MongoDBSession) Handler(w http.ResponseWriter, req *http.Request,
  next http.HandlerFunc, ctx domain.IContext) {
    s := session.Clone()
    defer s.Close()
    db := &MongoDB{
      currentDb: s.DB(session.DatabaseName),
    }
    SetMongoDbCtx(ctx, req, db)
    next(w, req)
}

func SetMongoDbCtx(ctx domain.IContext, r *http.Request, db *MongoDB) {
  ctx.Set(r, MongoDbKey, db)
}

func GetMongoDbCtx(ctx domain.IContext, r *http.Request) *MongoDB {
  if db := ctx.Get(r, MongoDbKey); db != nil {
    return db.(*MongoDB)
  }
  return nil
}


package domain

import (
  "net/http"
)

type ContextKey string

type IContext interface {
  Set(r *http.Request, key interface{}, val interface{})
  Get(r *http.Request, key interface{}) interface{}

  Inject(handler ContextHandlerFunc) http.HandlerFunc
}

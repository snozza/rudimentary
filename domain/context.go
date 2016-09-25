package domain

import (
  "net/http"
)

type ContextKey string

type IContext interface {
  Set(r *http.Request, key interface{}, val interface{})
  Get(r *http.Request, key interface{}) interface{}

  InjectMiddleware(ContextMiddlewareFunc) MiddlewareFunc
  Inject(handler ContextHandlerFunc) http.HandlerFunc
}

package context

import (
  "github.com/gorilla/context"
  "github.com/snozza/email-ads-data-api/domain"
  "net/http"
)

const DatabaseKey domain.ContextKey = "email-ads-data-api-mddlwr-context-database-key"

func New() *Context {
  return &Context{}
}

// Context implements IContext
type Context struct {
}

func (ctx *Context) InjectMiddleware(middleware domain.ContextMiddlewareFunc) domain.MiddlewareFunc {
  return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    middleware(rw, r, next, ctx)
  }
}

func (ctx *Context) Inject(handler domain.ContextHandlerFunc) http.HandlerFunc {
  return func(rw http.ResponseWriter, r *http.Request) {
    handler(rw, r, ctx)
  }
}

func (ctx *Context) Set(r *http.Request, key interface{}, val interface{}) {
  context.Set(r, key, val)
}

func (ctx *Context) Get(r *http.Request, key interface{}) interface{} {
  return context.Get(r, key)
}

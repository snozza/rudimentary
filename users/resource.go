package users

import (
  . "github.com/snozza/email-ads-data-api/users/domain"

  "github.com/snozza/email-ads-data-api/domain"
  "net/http"
)

type Options struct {
  BasePath string
  Database domain.IDatabase
  Renderer domain.IRenderer
  UserRepositoryFactory IUserRepositoryFactory
}

func NewResource(ctx domain.IContext, options *Options) *Resource {
  database := options.Database
  if database == nil {
    panic("users.Options.Database is required")
  }
  renderer := options.Renderer
  if renderer == nil {
    panic("users.Options.Renderer is required")
  }

  userRepositoryFactory := options.UserRepositoryFactory
  if userRepositoryFactory == nil {
    // init default UserRepositoryFactory
    userRepositoryFactory = NewUserRepositoryFactory()
  }

  u := &Resource{
    ctx,
    options,
    nil,
    database,
    renderer,
    userRepositoryFactory,
  }
  u.generateRoutes(options.BasePath)
  return u
}

type Resource struct {
  ctx domain.IContext
  options *Options
  routes *domain.Routes
  Database domain.IDatabase
  Renderer domain.IRenderer
  UserRepositoryFactory IUserRepositoryFactory
}

func (resource *Resource) Context() domain.IContext {
  return resource.ctx
}

func (resource *Resource) Routes() *domain.Routes {
  return resource.routes
}

func (resource *Resource) Render(w http.ResponseWriter, req *http.Request, status int, v interface{}) {
  resource.Renderer.Render(w, req, status, v)
}

func (resource *Resource) UserRepository(req *http.Request) IUserRepository {
  return resource.UserRepositoryFactory.New(resource.Database)
}

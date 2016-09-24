package server

import (
  "errors"
  "fmt"
  "github.com/gorilla/mux"
  "github.com/snozza/rudimentary/domain"
  "net/http"
)

// Router type
type Router struct {
  *mux.Router
  ctx domain.IContext
}

// TODO matcherFunc matches the handler to the correct API version based on its 'accept' header
// TODO: refactor matcher function as server.Config
func matchHandler(r domain.Route, defaultHandler http.HandlerFunc, ctx domain.IContext)
  func(r *http.Request, rm *mux.RouteMatch) bool {
    return func(req *http.Request, rm *mux.RouteMatch) bool {
      rm.Handler = defaultHandler
      return true
    }
}

// NewRouter
func NewRouter(ctx domain.IContext) *Router {
  router := mux.NewRouter().StrictSlash(true)
  return &Router{router, ctx}
}

func (router *Router) AddRoutes(routes *domain.Routes) *Router {
  if routes == nil {
    return Router
  }
  for _, route := range *routes {
    // get the defaultHandler for current route at init time so that we can safely panic
    // if it was not  defined
    defaultHandler, ok := route.RouteHandlers[route.DefaultVersion]
    if !ok {
      // server/router instantiation error
      // its safe to throw panic here
      panic(errors.New(fmt.Springf("Routes definition error, missing default route handler
        for version `%v`in `%v`", route.DefaultVersion, route.Name)))
    }
    router
      .Methods(route.Method)
      .Path(route.Pattern)
      .Name(route.Name)
      .MatcherFunc(matcherFunc(route, defaultHandler, router.ctx))
  }
  return router
}

func (router *Router) AddResources(resources ...domain.IResource) *Router {
  for _, resource := range resources {
    if resource.Routes() == nil {
      // server/routes instantiation error
      // its safe to panic here
      panic(errors.New(fmt.Sprintf("Routes definition missing: `%v`", resource)))
    }
    router.AddRoutes(resource.Routes())
  }
  return router
}

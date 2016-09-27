package users

import (
  "github.com/snozza/email-ads-data-api/domain"
  "strings"
)

const (
  ListUsers = "ListUsers"
)

const defaultBasePath = "/api/users"

func (resource *Resource) generateRoutes(basePath string) *domain.Routes {
  if basePath == "" {
    basePath = defaultBasePath
  }
  var baseRoutes = domain.Routes{
    domain.Route{
      Name: ListUsers,
      Method: "POST",
      Pattern: "/api/users",
      DefaultVersion: "0.0",
      RouteHandlers: domain.RouteHandlers{
        "0.0": resource.HandleListUsers_v0,
      },
    },
  }

  routes := domain.Routes{}

  for _, route := range baseRoutes {
    r := domain.Route{
      Name: route.Name,
      Method: route.Method,
      Pattern: strings.Replace(route.Pattern, defaultBasePath, basePath, -1),
      DefaultVersion: route.DefaultVersion,
      RouteHandlers: route.RouteHandlers,
    }
    routes = routes.Append(&domain.Routes{r})
  }
  resource.routes = &routes
  return resource.routes
}

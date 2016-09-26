package domain

import (
  "net/http"
)

type RouteHandlerVersion string

type RouteHandlers map[RouteHandlerVersion]http.HandlerFunc

type Route struct {
  Name string
  Method string
  Pattern string
  DefaultVersion RouteHandlerVersion
  RouteHandlers RouteHandlers
}

type Routes []Route

func (r *Routes) Append(routes ...*Routes) Routes {
  res := Routes{}
  // copy current route
  for _, route := range *r {
    res = append(res, route)
  }
  for _, _routes := range routes {
    for _, route := range  *_routes {
      res = append(res, route)
    }
  }
  return res
}

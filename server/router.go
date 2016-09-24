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

// matcherFunc matches the handler to the correct API version based on its 'accept' header
// TODO: refactor matcher function as server.Config



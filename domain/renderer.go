package domain

import (
  "net/http"
)

// Renderer interface
// TODO Add Data, XML
type IRenderer interface {
  Render(w http.ResponseWriter, req *http.Request, status int, v interface{})
  JSON(w http.ResponseWriter, status int, v interface{})
  Text(w http.ResponseWriter, status int, v []byte)
}

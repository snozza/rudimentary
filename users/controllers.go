package users

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
)

//------ User Request API v0 -------

type ListUsersResponse_v0 struct {
  Users Users `json:"users"`
  LastID string `json:last_id, omitempty"`
  Message string `json:"message, omitempty"`
  Success bool `json:"success"`
}

type ErrorResponse_v0 struct {
  Message string `json:"message",omitempty"`
  Success bool `json:"success"`
}

func (resource *Resource) DecodeRequestBody(w http.ResponseWriter, req *http.Request, target interface{}) error {
  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(target)
  if err != nil {
    resource.RenderError(w, req, http.StatusBadRequest,
      fmt.Sprintf("Request body parse error: %v", err.Error()))
    return err
  }
  return nil
}

func (resource *Resource) RenderError(w http.ResponseWriter, req *http.Request, status int, message string) {
  resource.Render(w, req, status, ErrorResponse_v0{
    Message: message,
    Success: false,
  })
}

func (resource *Resource) HandleListUsers_v0(w http.ResponseWriter, req *http.Request) {
  repo := resource.UserRepository(req)

  // filter & pagination params
  field := req.FormValue("field")
  query := req.FormValue("q")
  lastID := req.FormValue("last_id")
  perPageStr := req.FormValue("per_page")
  sort := req.FormValue("sort")

  perPage, err := strconv.Atoi(perPageStr)
  if err != nil {
    perPage = 20
  }

  u := repo.FilterUsers(field, query, lastID, perPage, sort)
  users := *u.(*Users)
  if len(users) > 0 {
    lastID = users[len(users) - 1].ID.Hex()
  }
  resource.Render(w, req, http.StatusOK, ListUsersResponse_v0{
    Users: users,
    LastID: lastID,
    Message: "User list retrieved",
    Success: true,
  })
}

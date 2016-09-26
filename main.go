package main

import (
  "fmt"
  "time"
  "net/http"
  "github.com/snozza/rudimentary/middleware/context"
  "github.com/snozza/rudimentary/middleware/mongodb"
  "github.com/snozza/rudimentary/middleware/renderer"
  "github.com/snozza/rudimentary/users"
  "github.com/snozza/rudimentary/server"
)

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path);
  fmt.Fprintf(w, "Welcome, %1", r.URL.Path[1:])
}

func main() {

  ctx := context.New()

  db := mongodb.New(&mongodb.Options{
    ServerName: "localhost",
    DatabaseName: "rudimentary",
  })

  _ = db.NewSession()

  // init server
  s := server.NewServer(&server.Config{
      Context: ctx,
  })

  renderer := renderer.New(&renderer.Options{
    IndentJSON: true,
  }, renderer.JSON)

  // set up users resource
  usersResource := users.NewResource(ctx, &users.Options{
    Database: db,
    Renderer: renderer,
  })

  // set up router
  router := server.NewRouter(s.Context)

  // add REST resources to router
  router.AddResources(usersResource)

  s.UseRouter(router)

  s.Run(":3001", server.Options{
    Timeout: 10 * time.Second,
  })
}

package main

import (
  "time"
  "github.com/snozza/email-ads-data-api/middleware/context"
  "github.com/snozza/email-ads-data-api/middleware/mongodb"
  "github.com/snozza/email-ads-data-api/middleware/renderer"
  "github.com/snozza/email-ads-data-api/users"
  "github.com/snozza/email-ads-data-api/server"
)

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

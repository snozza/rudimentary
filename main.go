package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "github.com/snozza/context"
  "github.com/snozza/mongodb"
)

type Message struct {
  Message string `json:"message"`
}

type Example struct {
  Name string
  Age int
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Println(r.URL.Path);
  fmt.Fprintf(w, "Welcome, %1", r.URL.Path[1:])
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
  m := Message{"Welcome to this terrible API, build v0.0.1"}
  b, err := json.Marshal(m)

  if err != nil {
    panic(err)
  }

  w.Write(b);
}

func main() {

  ctx := context.New()

  db := mongodb.New(&mongodb.Options{
    ServerName: "localhost",
    DatabaseName: "rudimentary"
  })

  // init server
  s := server.NewServer(&server.Config{
      Context: ctx
  })

  // set up router
  router := server.NewRouter(s.Context)

  s.UseRouter(router)

  s.Run(":3001", server.Options{
    Timeout 10*time.Second,
  })
}


  //c := session.DB("rudimentary").C("example")
  //err = c.Insert(&Example{"Andrew", 15},
                  //&Example{"Mike", 17})
  //if err != nil {
    //log.Fatal(err)
  //}

  //result := Example{}
  //err = c.Find(bson.M{"name": "Andrew"}).One(&result)
  //if err != nil {
    //log.Fatal(err)
  //}

  //fmt.Println("Age:", result.Age)

  //http.HandleFunc("/", index)
  //http.HandleFunc("/about", about)
  //http.ListenAndServe(":8080", nil)
//}

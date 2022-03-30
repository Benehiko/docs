package main

import (
  "fmt"
  ory "github.com/ory/client-go"
  "net/http"
  "os"
)

type App struct {
  ory *ory.APIClient
  // save the cookies for any upstream calls to the Ory apis
  cookies string
  // save the session to display it on the dashboard
  session *ory.Session
}

func main() {
  // we need to point to our custom domain now
  c := ory.NewConfiguration()
  c.Servers = ory.ServerConfigurations{{URL: "https://auth.example.com"}}
  app := &App{
    ory: ory.NewAPIClient(c),
  }
  mux := http.NewServeMux()

  // dashboard
  mux.Handle("/", app.sessionMiddleware(app.dashboardHandler()))

  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }

  fmt.Printf("Application launched and running on http://127.0.0.1:%s\n", port)
  // start the server
  http.ListenAndServe(":"+port, mux)
}

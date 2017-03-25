package main

import (
  "github.com/labstack/echo"
  "codefest/routes"
  "os"
)

func main() {
  port := os.Getenv("CODEFEST_PORT")
  if port == "" {
    port = "8000"
  }
  e := echo.New()
  routes.Init(e)
  e.Logger.Fatal(e.Start(":" + port))
}
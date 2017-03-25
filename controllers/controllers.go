package controllers

import (
  "github.com/labstack/echo"
  "net/http"
)

const OK = http.StatusOK

func Index (c echo.Context) error {
  return c.JSON(OK, "Index")
}

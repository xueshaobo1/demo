package controllers

import (
//    "net/http"
    "github.com/labstack/echo"
)

func Hello(c echo.Context) error {
  return c.Redirect(301, "/index")
}

func Index(c echo.Context) error {
  return c.Render(200, "index.html","")
}

package routes

import (
  "github.com/labstack/echo"
  "codefest/controllers"
)

func Init(e *echo.Echo) {
  e.GET("/", controllers.Index)

  codefest := e.Group("codefest/lectures")
  codefest.GET("", controllers.Likes)
  codefest.POST("/like", controllers.Like)
}

package router

import (
//    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "demo/controllers"
    "html/template"
)

func Run(){
    e := echo.New()
    //引用模板文件
    t := &Template{
    templates: template.Must(template.ParseGlob("view/*.html")),
    }
    e.Renderer = t
    //路由
    e.GET("/",controllers.Hello)
    e.GET("/index",controllers.Index)
    //添加中间件
    api := e.Group(`/api`, midJwt)
    api.GET("/sys",controllers.Index)


    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Logger.Fatal(e.Start(":1323"))

}

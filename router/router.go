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
    /*
    组可以单独写在其他配置文件中
    例如：
    router.go //写在router.go文件中
    api := e.Group(`/api`, midJwt)
    Api(api)
    
    api.go //写在api.go文件中
    package router
    import (
        "github.com/labstack/echo"
        "demo/controllers"
     )
    func Api(api *echo.Group){
    api.GET("/sys",controllers.Index)
    这样可以把路由拆开卸载不同的文件里面
    */
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Logger.Fatal(e.Start(":1323"))

}

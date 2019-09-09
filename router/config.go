package router

import (
    "html/template"
    "github.com/labstack/echo"
    "io"
)
//建立模板文件
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}
//建立中间件，方便判断登录状态
func midJwt(next echo.HandlerFunc) echo.HandlerFunc {
    return func(ctx echo.Context) error {
        ctx.Response().Header().Set(echo.HeaderServer, "dev")
        return next(ctx)
	}
}

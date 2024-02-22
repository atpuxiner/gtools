package addtpl

const appRouterTpl = `
package router

import (
	"github.com/atpuxiner/grapi/app/api/vn"
	"github.com/atpuxiner/grapi/app/middleware"
	"github.com/gin-gonic/gin"
)

func tplRouter() {
	registerVnRouter(func(rg *gin.RouterGroup) {
		tplApi := vn.NewTplApi()
		rgg := rg.Group("tpl") // 此处可添加公共中间件
		{
			rgg.GET("/:id", middleware.JwtAuth(), tplApi.Get)
		}
	})
}

func init() {
	tplRouter()
}
`

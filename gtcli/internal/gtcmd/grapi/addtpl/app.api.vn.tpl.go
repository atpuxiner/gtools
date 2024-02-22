package addtpl

const appApiVnTpl = `
package vn

import (
	"github.com/atpuxiner/grapi/app/api"
	"github.com/atpuxiner/grapi/app/business/tpl"
	"github.com/atpuxiner/grapi/app/utils/status"
	"github.com/gin-gonic/gin"
)

type TplApi struct {
	api.BaseApi
}

func NewTplApi() TplApi {
	return TplApi{}
}

var tplBusiness = tpl.NewBusiness()

// @Tags Api.Tpl
// @Summary 查询tpl详情
// @Description 查询tpl详情
// @Param id path uint true "业务id"
// @Success 0 {object} api.ResponseJson "请求成功"
// @Failure 1 {object} api.ResponseJson "请求失败"
// @Router /api/vn/tpl/{id} [get]
func (r TplApi) Get(c *gin.Context) {
	// 参数处理
	id, err := r.GetParamUint(c, "id")
	if err != nil {
		r.Failure(c, status.CodeBadParam, err)
		return
	}
	// 数据处理
	data, err := tplBusiness.Get(id)
	r.Response(c, data, err)
}
`

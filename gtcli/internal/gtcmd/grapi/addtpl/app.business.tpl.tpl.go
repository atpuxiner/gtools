package addtpl

const appBusinessTplTpl = `
package tpl

import (
	"github.com/atpuxiner/grapi/app/datatype/entity/enttpl"
	"github.com/atpuxiner/grapi/app/initializer/db"
)

type Business struct{}

func NewBusiness() Business {
	return Business{}
}

func (r Business) Get(id uint) (any, error) {
	var tpl enttpl.GetResp
	if err := db.DB.First(&tpl, id).Error; err != nil {
		return nil, err
	}
	return tpl, nil
}
`

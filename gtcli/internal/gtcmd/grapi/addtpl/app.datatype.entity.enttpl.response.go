package addtpl

const appDatatypeEntityEnttplResponse = `
package enttpl

import (
	"github.com/atpuxiner/grapi/app/datatype/entity"
	"github.com/atpuxiner/grapi/app/datatype/model"
)

type GetResp struct {
	model.BaseTpl
	entity.DefaultEntityWithoutDlt
}
`

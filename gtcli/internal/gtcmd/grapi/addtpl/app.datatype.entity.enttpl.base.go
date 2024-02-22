package addtpl

const appDatatypeEntityEnttplBase = `
package enttpl

import (
	"github.com/atpuxiner/grapi/app/datatype/entity"
	"github.com/atpuxiner/grapi/app/datatype/model"
)

type Tpl struct {
	model.BaseTpl
	entity.DefaultEntity
}
`

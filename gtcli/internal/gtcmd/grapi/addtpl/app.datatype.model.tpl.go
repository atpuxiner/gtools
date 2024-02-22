package addtpl

const appDatatypeModelTpl = `
package model

const TblNameTpl = "tpl"

type BaseTpl struct{}

func (BaseTpl) TableName() string {
	return TblNameTpl
}

type Tpl struct {
	BaseTpl
	DefaultModel
}

func init() {
	registerModel(&Tpl{})
}
`

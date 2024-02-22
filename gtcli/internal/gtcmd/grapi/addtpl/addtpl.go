package addtpl

import "path/filepath"

var Tpls = map[string]string{
	filepath.FromSlash("app/api/vn/tpl.go"):                      appApiVnTpl,
	filepath.FromSlash("app/business/tpl/tpl.go"):                appBusinessTplTpl,
	filepath.FromSlash("app/datatype/entity/enttpl/base.go"):     appDatatypeEntityEnttplBase,
	filepath.FromSlash("app/datatype/entity/enttpl/request.go"):  appDatatypeEntityEnttplRequest,
	filepath.FromSlash("app/datatype/entity/enttpl/response.go"): appDatatypeEntityEnttplResponse,
	filepath.FromSlash("app/datatype/entity/enttpl/validate.go"): appDatatypeEntityEnttplValidate,
	filepath.FromSlash("app/datatype/model/tpl.go"):              appDatatypeModelTpl,
	filepath.FromSlash("app/router/tpl.go"):                      appRouterTpl,
}

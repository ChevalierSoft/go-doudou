package codegen

import (
	"github.com/iancoleman/strcase"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/astutils"
)

func genHttp(dir string, ic astutils.InterfaceCollector) {
	GenConfig(dir)
	GenHttpMiddleware(dir)
	GenHttpHandler(dir, ic, 0)
	GenHttpHandlerImpl(dir, ic, GenHttpHandlerImplConfig{
		CaseConvertor: strcase.ToLowerCamel,
	})
	GenSvcImpl(dir, ic)
	GenDoc(dir, ic, GenDocConfig{})
}

package client

import (
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/v2/cmd/internal/openapi/v3/codegen"
	v3 "github.com/unionj-cloud/go-doudou/v2/toolkit/openapi/v3"
	"os"
	"path/filepath"
	"strings"
)

var dtoTmpl = `package {{.Pkg}}

{{- range $k, $v := .Schemas }}
{{ toComment $v.Description ($k | toCamel)}}
type {{$k | toCamel}} struct {
{{- range $pk, $pv := $v.Properties }}
	{{ $pv.Description | toComment }}
	{{- if stringContains $v.Required $pk }}
	// required
	{{ $pk | toCamel}} {{$pv | toGoType }} ` + "`" + `json:"{{$pk}}{{if $.Omit}},omitempty{{end}}" url:"{{$pk}}"` + "`" + `
	{{- else }}
	{{ $pk | toCamel}} {{$pv | toOptionalGoType }} ` + "`" + `json:"{{$pk}}{{if $.Omit}},omitempty{{end}}" url:"{{$pk}}"` + "`" + `
	{{- end }}
{{- end }}
}
{{- end }}
`

// GenGoClient generate go http client code from OpenAPI3.0 json document
func GenGoClient(dir string, file string, omit bool, env, pkg string) {
	var (
		err       error
		f         *os.File
		clientDir string
		fi        os.FileInfo
		api       v3.API
		dtoFile   string
	)
	clientDir = filepath.Join(dir, pkg)
	if err = os.MkdirAll(clientDir, os.ModePerm); err != nil {
		panic(err)
	}
	api = v3.LoadAPI(file)
	generator := &codegen.OpenAPICodeGenerator{
		Schemas:       api.Components.Schemas,
		RequestBodies: api.Components.RequestBodies,
		Responses:     api.Components.Responses,
		Omitempty:     omit,
	}
	svcmap := make(map[string]map[string]v3.Path)
	for endpoint, path := range api.Paths {
		svcname := strings.Split(strings.Trim(endpoint, "/"), "/")[0]
		if value, exists := svcmap[svcname]; exists {
			value[endpoint] = path
		} else {
			svcmap[svcname] = make(map[string]v3.Path)
			svcmap[svcname][endpoint] = path
		}
	}
	for svcName, paths := range svcmap {
		generator.GenGoHTTP(paths, svcName, clientDir, env, pkg)
	}
	dtoFile = filepath.Join(clientDir, "dto.go")
	fi, err = os.Stat(dtoFile)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	if fi != nil {
		logrus.Warningln("file dto.go will be overwritten")
	}
	if f, err = os.Create(dtoFile); err != nil {
		panic(err)
	}
	defer f.Close()
	generator.GenGoDto(api.Components.Schemas, dtoFile, pkg, dtoTmpl)
}

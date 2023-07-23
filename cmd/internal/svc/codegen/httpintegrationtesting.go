package codegen

import (
	"bytes"
	"fmt"
	"github.com/rbretecher/go-postman-collection"
	"github.com/unionj-cloud/go-doudou/v2/cmd/internal/templates"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/astutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"github.com/unionj-cloud/go-doudou/v2/version"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var appendIntegrationTestingTmpl = `
{{- range $response := .Responses }}

func Test_{{$response.Name | cleanName}}(t *testing.T) {
	apitest.New("{{$response.Name}}").
		Handler(router).
		{{$response.OriginalRequest.Method | toString | capital}}("{{ $response.OriginalRequest.URL.Path | toEndpoint }}").
{{- range $header := $response.OriginalRequest.Header }}
{{- if not $header.Disabled }}
		Header("{{$header.Key}}", "{{$header.Value}}").
{{- end }}
{{- end }}
{{- if $response.OriginalRequest.URL.Query }}
{{- range $query := $response.OriginalRequest.URL.Query }}
{{- if not (index $query "disabled") }}
		Query("{{index $query "key"}}", "{{index $query "value"}}").
{{- end }}
{{- end }}
{{- end }}
{{- if $response.OriginalRequest.Body }}
{{- if eq $response.OriginalRequest.Body.Mode "raw" }}
		JSON(` + "`" + `{{$response.OriginalRequest.Body.Raw}}` + "`" + `).
{{- else if eq $response.OriginalRequest.Body.Mode "urlencoded" }}
{{- range $query := $response.OriginalRequest.Body.URLEncoded }}
{{- if not (index $query "disabled") }}
		FormData("{{index $query "key"}}", "{{index $query "value"}}").
{{- end }}
{{- end }}
{{- else if eq $response.OriginalRequest.Body.Mode "formdata" }}
{{- range $query := $response.OriginalRequest.Body.FormData }}
{{- if not (index $query "disabled") }}
{{- if eq (index $query "type") "file" }}
		MultipartFile("{{index $query "key"}}", "{{index $query "src"}}").
{{- else }}
		MultipartFormData("{{index $query "key"}}", "{{index $query "value"}}").
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
		Expect(t).
{{- if $response.Body }}
		Body(` + "`" + `{{$response.Body}}` + "`" + `).
{{- end }}
		Status({{$response.Code}}).
		End()
}
{{- end }}
`

var integrationTestingImportTmpl = `
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"net/http"
	{{.ServiceAlias}} "{{.ServicePackage}}"
	"{{.ServicePackage}}/config"
	"{{.ServicePackage}}/db"
	"{{.ServicePackage}}/transport/httpsrv"
	"testing"
`

var initIntegrationTestingTmpl = templates.EditableHeaderTmpl + `package integrationtest_test

import ()

var (
	router *mux.Router
)

func TestMain(m *testing.M) {
	_ = godotenv.Load("{{.DotenvPath}}")
	conf := config.LoadFromEnv()
	conn, err := db.NewDb(conf.DbConf)
	if err != nil {
		panic(err)
	}
	defer func() {
		if conn == nil {
			return
		}
		if err := conn.Close(); err == nil {
			zlogger.Info().Msg("Database connection is closed")
		} else {
			zlogger.Warn().Msg("Failed to close database connection")
		}
	}()
	svc := {{.ServiceAlias}}.New{{.SvcName}}(conf, conn)
	handler := httpsrv.New{{.SvcName}}Handler(svc)
	router = mux.NewRouter()
	for _, item := range httpsrv.Routes(handler) {
		router.
			Methods(item.Method).
			Path(item.Pattern).
			Name(item.Name).
			Handler(item.HandlerFunc)
	}
	m.Run()
}
` + appendIntegrationTestingTmpl

func toEndpoint(input []string) string {
	return "/" + strings.Join(input, "/")
}

func toString(input postman.Method) string {
	return string(input)
}

func capital(input string) string {
	return strings.Title(strings.ToLower(input))
}

func GenHttpIntegrationTesting(dir string, ic astutils.InterfaceCollector, postmanCollectionPath, dotenvPath string) {
	var (
		err                error
		testFile           string
		f                  *os.File
		tpl                *template.Template
		buf                bytes.Buffer
		fi                 os.FileInfo
		tmpl               string
		importBuf          bytes.Buffer
		integrationTestDir string
		responses          []*postman.Response
	)
	integrationTestDir = filepath.Join(dir, "integrationtest")
	if err = os.MkdirAll(integrationTestDir, os.ModePerm); err != nil {
		panic(err)
	}
	testFile = filepath.Join(integrationTestDir, "integration_test.go")
	responses = notGenerated(integrationTestDir, postmanCollectionPath)
	fi, _ = os.Stat(testFile)
	if fi != nil {
		zlogger.Warn().Msg("New content will be append to integration_test.go file")
		if f, err = os.OpenFile(testFile, os.O_APPEND, os.ModePerm); err != nil {
			panic(err)
		}
		defer f.Close()
		tmpl = appendIntegrationTestingTmpl
	} else {
		if f, err = os.Create(testFile); err != nil {
			panic(err)
		}
		defer f.Close()
		tmpl = initIntegrationTestingTmpl
	}

	servicePkg := astutils.GetPkgPath(dir)

	funcMap := make(map[string]interface{})
	funcMap["toEndpoint"] = toEndpoint
	funcMap["toString"] = toString
	funcMap["capital"] = capital
	funcMap["cleanName"] = cleanName
	if tpl, err = template.New("integration_test.go.tmpl").Funcs(funcMap).Parse(tmpl); err != nil {
		panic(err)
	}
	absDotenv, _ := filepath.Abs(dotenvPath)
	absDir, _ := filepath.Abs(integrationTestDir)
	relDotenv, _ := filepath.Rel(absDir, absDotenv)
	if err = tpl.Execute(&buf, struct {
		ServicePackage string
		ServiceAlias   string
		Version        string
		DotenvPath     string
		SvcName        string
		Responses      []*postman.Response
	}{
		ServicePackage: servicePkg,
		ServiceAlias:   ic.Package.Name,
		Version:        version.Release,
		DotenvPath:     relDotenv,
		SvcName:        ic.Interfaces[0].Name,
		Responses:      responses,
	}); err != nil {
		panic(err)
	}
	original, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	original = append(original, buf.Bytes()...)
	if tpl, err = template.New("testimportimpl.go.tmpl").Parse(integrationTestingImportTmpl); err != nil {
		panic(err)
	}
	if err = tpl.Execute(&importBuf, struct {
		ServicePackage string
		ServiceAlias   string
	}{
		ServicePackage: servicePkg,
		ServiceAlias:   ic.Package.Name,
	}); err != nil {
		panic(err)
	}
	original = astutils.AppendImportStatements(original, importBuf.Bytes())
	astutils.FixImport(original, testFile)
}

func notGenerated(dir, postmanCollectionPath string) []*postman.Response {
	file, err := os.Open(postmanCollectionPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	c, err := postman.ParseCollection(file)
	if err != nil {
		panic(err)
	}
	responses := flattenResponses(c.Items)
	var files []string
	err = filepath.Walk(dir, astutils.Visit(&files))
	if err != nil {
		panic(err)
	}
	sc := astutils.NewStaticMethodCollector(astutils.ExprString)
	for _, file := range files {
		root, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
		ast.Walk(sc, root)
	}
	methodStore := make(map[string]struct{})
	for _, item := range sc.Methods {
		methodStore[item.Name] = struct{}{}
	}
	var result []*postman.Response
	for _, item := range responses {
		methodName := fmt.Sprintf("Test_%s", cleanName(item.Name))
		if _, exists := methodStore[methodName]; !exists {
			result = append(result, item)
		}
	}
	return result
}

func flattenResponses(items []*postman.Items) []*postman.Response {
	var result []*postman.Response
	for _, item := range items {
		if len(item.Items) > 0 {
			result = append(result, flattenResponses(item.Items)...)
		} else {
			for _, resp := range item.Responses {
				if resp.Name != "Untitled Example" && resp.Name != "response" {
					result = append(result, resp)
				}
			}
		}
	}
	return result
}

func cleanName(name string) string {
	name = strings.ReplaceAll(strings.ReplaceAll(name, "{", ""), "}", "")
	name = strings.ReplaceAll(strings.Trim(name, "/"), "/", "_")
	return name
}

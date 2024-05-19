package rest

import (
	"bytes"
	"fmt"
	"github.com/rs/cors"
	"html/template"
	"net/http"
	"strings"
)

//	 window.docs = [
//	{
//	  "label": "Banana",
//	  "value": "http://localhost:6060/banana/go-doudou/doc",
//	},
//	{
//	  "label": "Apple",
//	  "value": "http://localhost:6060/apple/go-doudou/doc",
//	}
//
// ]
var Docs []DocItem
var DocRoutes = docRoutes

type DocItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func docRoutes(doc string) []Route {
	routes := []Route{
		{
			Name:    "GetDoc",
			Method:  "GET",
			Pattern: "/go-doudou/doc",
			HandlerFunc: func(_writer http.ResponseWriter, _req *http.Request) {
				var (
					tpl *template.Template
					err error
					buf bytes.Buffer
				)
				if tpl, err = template.New("onlinedoc.tmpl").Parse(onlinedocTmpl); err != nil {
					panic(err)
				}
				docUrl := "openapi.json"
				if err = tpl.Execute(&buf, struct {
					Doc    string
					DocUrl string
					Docs   []DocItem
				}{
					Doc:    doc,
					DocUrl: docUrl,
					Docs:   Docs,
				}); err != nil {
					panic(err)
				}
				_writer.Header().Set("Content-Type", "text/html; charset=utf-8")
				_writer.Write(buf.Bytes())
			},
		},
		{
			Name:    "GetOpenAPI",
			Method:  "GET",
			Pattern: "/go-doudou/openapi.json",
			HandlerFunc: func(_writer http.ResponseWriter, _req *http.Request) {
				_writer.Write([]byte(doc))
			},
		},
	}
	corsOpts := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*",
		},

		AllowOriginRequestFunc: func(r *http.Request, origin string) bool {
			if strings.Contains(r.URL.Path, fmt.Sprintf("%sopenapi.json", gddPathPrefix)) {
				return true
			}
			return false
		},
	})
	gddmiddlewares := []MiddlewareFunc{corsOpts.Handler, MiddlewareFunc(basicAuth())}

	for k, item := range routes {
		h := http.Handler(item.HandlerFunc)
		for i := len(gddmiddlewares) - 1; i >= 0; i-- {
			h = gddmiddlewares[i].Middleware(h)
		}
		item.HandlerFunc = h.(http.HandlerFunc)
		routes[k] = item
	}

	return routes
}

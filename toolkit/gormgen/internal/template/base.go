package template

const NotEditMark = `
// Code generated by github.com/wubin1989/gen. DO NOT EDIT.
// Code generated by github.com/wubin1989/gen. DO NOT EDIT.
// Code generated by github.com/wubin1989/gen. DO NOT EDIT.
`

const EditMark = `
// Code generated by github.com/wubin1989/gen. YOU CAN EDIT.
// Code generated by github.com/wubin1989/gen. YOU CAN EDIT.
// Code generated by github.com/wubin1989/gen. YOU CAN EDIT.
`

const NotEditMarkForGDDShort = `// Code generated by github.com/wubin1989/gen for go-doudou. DO NOT EDIT.`

const EditMarkForGDD = `
// Code generated by github.com/wubin1989/gen for go-doudou. YOU CAN EDIT.
// Code generated by github.com/wubin1989/gen for go-doudou. YOU CAN EDIT.
// Code generated by github.com/wubin1989/gen for go-doudou. YOU CAN EDIT.
`

const Header = NotEditMark + `
package {{.Package}}

import(	
	{{range .ImportPkgPaths}}{{.}}` + "\n" + `{{end}}
)
`

package codegen

import (
	"bufio"
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/cmd/internal/astutils"
	v3 "github.com/unionj-cloud/go-doudou/cmd/internal/protobuf/v3"
	"github.com/unionj-cloud/go-doudou/version"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var mainTmplGrpc = `/**
* Generated by go-doudou {{.Version}}.
* You can edit it as your need.
*/
package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/unionj-cloud/go-doudou/toolkit/zlogger"
	"google.golang.org/grpc"
	ddgrpc "github.com/unionj-cloud/go-doudou/framework/grpc"
	{{.ServiceAlias}} "{{.ServicePackage}}"
    "{{.ConfigPackage}}"
	pb "{{.PbPackage}}"
)

func main() {
	conf := config.LoadFromEnv()
	svc := {{.ServiceAlias}}.New{{.SvcName}}(conf)
	grpcServer := ddgrpc.NewGrpcServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			tags.StreamServerInterceptor(tags.WithFieldExtractor(tags.CodeGenRequestFieldExtractor)),
			logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			tags.UnaryServerInterceptor(tags.WithFieldExtractor(tags.CodeGenRequestFieldExtractor)),
			logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(zlogger.Logger)),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.Register{{.GrpcSvcName}}Server(grpcServer, svc)
	grpcServer.Run()
}
`

// GenMainGrpc generates main function for grpc service
func GenMainGrpc(dir string, ic astutils.InterfaceCollector, grpcSvc v3.Service) {
	var (
		err       error
		modfile   string
		modName   string
		mainfile  string
		firstLine string
		f         *os.File
		tpl       *template.Template
		cmdDir    string
		svcName   string
		alias     string
		sqlBuf    bytes.Buffer
		source    string
	)
	cmdDir = filepath.Join(dir, "cmd")
	if err = MkdirAll(cmdDir, os.ModePerm); err != nil {
		panic(err)
	}

	svcName = ic.Interfaces[0].Name
	alias = ic.Package.Name
	mainfile = filepath.Join(cmdDir, "main.go")
	if _, err = Stat(mainfile); os.IsNotExist(err) {
		modfile = filepath.Join(dir, "go.mod")
		if f, err = Open(modfile); err != nil {
			panic(err)
		}
		reader := bufio.NewReader(f)
		firstLine, _ = reader.ReadString('\n')
		modName = strings.TrimSpace(strings.TrimPrefix(firstLine, "module"))

		if f, err = Create(mainfile); err != nil {
			panic(err)
		}
		defer f.Close()

		if tpl, err = template.New("main.go.tmpl").Parse(mainTmplGrpc); err != nil {
			panic(err)
		}
		if err = tpl.Execute(&sqlBuf, struct {
			ServicePackage string
			ConfigPackage  string
			PbPackage      string
			SvcName        string
			ServiceAlias   string
			Version        string
			GrpcSvcName    string
		}{
			ServicePackage: modName,
			ConfigPackage:  modName + "/config",
			PbPackage:      modName + "/transport/grpc",
			SvcName:        svcName,
			ServiceAlias:   alias,
			Version:        version.Release,
			GrpcSvcName:    grpcSvc.Name,
		}); err != nil {
			panic(err)
		}
		source = strings.TrimSpace(sqlBuf.String())
		astutils.FixImport([]byte(source), mainfile)
	} else {
		logrus.Warnf("file %s already exists", mainfile)
	}
}

package codegen

import (
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/v2/cmd/internal/templates"
	"github.com/unionj-cloud/go-doudou/v2/version"
	"os"
	"path/filepath"
	"text/template"
)

var configTmpl = templates.EditableHeaderTmpl + `package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/errorx"
)

type Config struct {
	BizConf BizConf
}

type BizConf struct {
	ApiSecret string ` + "`" + `split_words:"true"` + "`" + `
}

func LoadFromEnv() *Config {
	var bizConf BizConf
	err := envconfig.Process("biz", &bizConf)
	if err != nil {
		errorx.Panic("Error processing environment variables")
	}
	return &Config{
		BizConf: bizConf,
	}
}
`

//GenConfig generates config file
func GenConfig(dir string) {
	var (
		err        error
		configfile string
		f          *os.File
		tpl        *template.Template
		configDir  string
	)
	configDir = filepath.Join(dir, "config")
	if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
		panic(err)
	}

	configfile = filepath.Join(configDir, "config.go")
	if _, err = os.Stat(configfile); os.IsNotExist(err) {
		if f, err = os.Create(configfile); err != nil {
			panic(err)
		}
		defer f.Close()
		tpl, _ = template.New("config.go.tmpl").Parse(configTmpl)
		_ = tpl.Execute(f, struct {
			Version string
		}{
			Version: version.Release,
		})
	} else {
		logrus.Warnf("file %s already exists", configfile)
	}
}

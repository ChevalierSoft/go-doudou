package ddhttp

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/svc/config"
	"github.com/unionj-cloud/go-doudou/svc/registry"
)

// Seed starts a seed node
func Seed() {
	config.GddServiceName.Write("seed")
	config.GddPort.Write("56200")
	config.GddMemPort.Write("56199")
	err := registry.NewNode()
	if err != nil {
		logrus.Panicln(fmt.Sprintf("%+v", err))
	}
	srv := NewDefaultHttpSrv()
	srv.Run()
}

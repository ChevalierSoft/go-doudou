package ddhttp

import (
	"fmt"
	"github.com/unionj-cloud/go-doudou/svc/config"
	"github.com/unionj-cloud/go-doudou/svc/logger"
	"github.com/unionj-cloud/go-doudou/svc/registry"
)

// Seed starts a seed node
func Seed() {
	config.GddServiceName.Write("seed")
	err := registry.NewNode()
	if err != nil {
		logger.Panicln(fmt.Sprintf("%+v", err))
	}
	srv := NewDefaultHttpSrv()
	srv.Run()
}

package registry

import (
	"github.com/unionj-cloud/go-doudou/framework/http/model"
	"net/http"
)

// ConfigHandler define http handler interface
type ConfigHandler interface {
	GetConfig(w http.ResponseWriter, r *http.Request)
}

// Routes return route slice for gorilla mux
func Routes() []model.Route {
	handler := NewConfigHandler()
	return []model.Route{
		{
			Name:        "GetConfig",
			Method:      "GET",
			Pattern:     "/go-doudou/config",
			HandlerFunc: handler.GetConfig,
		},
	}
}

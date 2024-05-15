package router

import (
	"github.com/gorilla/mux"
	api_controller "github.com/ternaryinvalid/safenet/server/internal/app/adapters/primary/http-adapter/api-controller"
	"net/http"
)

type Router struct {
	router *mux.Router
}

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.Handler
}

func New() *Router {
	router := mux.NewRouter()

	r := Router{
		router: router,
	}

	return &r
}

func (r *Router) Router() http.Handler {
	return r.router
}

const (
	apiV1Prefix = "/api/v1"
)

func (r *Router) AppendRoutesV1(controller *api_controller.ApiController) {
	apiV1Subrouter := r.router.PathPrefix(apiV1Prefix).Subrouter()

	routes := []Route{
		{
			Name:    "get messages",
			Path:    "/messages",
			Method:  http.MethodPost,
			Handler: http.HandlerFunc(controller.GetMessages),
		},
		{
			Name:    "send message",
			Path:    "/send",
			Method:  http.MethodPost,
			Handler: http.HandlerFunc(controller.SendMessage),
		},
		{
			Name:    "generate keys",
			Path:    "/generate-keys",
			Method:  http.MethodGet,
			Handler: http.HandlerFunc(controller.GenerateKeys),
		},
	}

	r.appendRoutesToRouter(apiV1Subrouter, routes)
}

func (r *Router) appendRoutesToRouter(subrouter *mux.Router, routes []Route) {
	for _, route := range routes {
		subrouter.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}

package api_controller

import "github.com/ternaryinvalid/safenet/client/internal/app/application"

type ApiController struct {
	app *application.Application
}

func New(app *application.Application) *ApiController {
	return &ApiController{
		app: app,
	}
}

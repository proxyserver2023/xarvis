package router

import "github.com/alamin-mahamud/xarvis/pkg/handler"

func getAuthRoutes() Routes {

	userHandler := handler.NewUser()
	return Routes{
		Route{"List", "GET", "", userHandler.List},
		Route{"Get", "GET", "/{id}", userHandler.Get},

		Route{"Create", "POST", "", userHandler.Create},
		Route{"Update", "PATCH", "/{id}", userHandler.Update},
		Route{"Delete", "DELETE", "/{id}", userHandler.Delete},
	}
}

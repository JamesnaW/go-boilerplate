package route

import (
	"github.com/timeff/go-boilerplate/handlers"
	"github.com/timeff/go-boilerplate/models"
)

func Todo() model.Routes {

	return model.Routes{
		model.Route{
			"TodoIndex",
			"GET",
			"/",
			handlers.TodoIndex,
		},
		model.Route{
			"TodoShow",
			"GET",
			"/todos/:param",
			handlers.TodoShow,
		},
	}
}

package route

import (
	"github.com/JamesnaW/boilerplate/handlers"
	"github.com/JamesnaW/boilerplate/models"
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

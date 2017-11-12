package route

import (
	"github.com/JamesnaW/go-boilerplate/handlers"
	"github.com/JamesnaW/go-boilerplate/models"
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

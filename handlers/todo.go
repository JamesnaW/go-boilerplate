package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"go-boilerplate/models"

	"github.com/valyala/fasthttp"
)

func TodoIndex(ctx *fasthttp.RequestCtx) {
	todo := model.Todo{Name: "Write presentation"}
	todo.Complete()
	todos := model.Todos{
		todo,
		model.Todo{Name: "Host meetup", Due: time.Now()},
	}
	json.NewEncoder(ctx).Encode(todos)
}

func TodoShow(ctx *fasthttp.RequestCtx) {
	todoID := ctx.UserValue("param")
	query := ctx.FormValue("q")
	fmt.Fprintf(ctx, "Todo show: %s", todoID)

	if query != nil {
		fmt.Fprintf(ctx, ", Query string: %s", string(query))
	}
}

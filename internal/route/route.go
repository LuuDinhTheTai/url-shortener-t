package route

import (
	"url-shortener-t/internal/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var routes = []Route{
	{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
}

func NewRouter(r *gin.Engine) {
	for _, route := range routes {
		r.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

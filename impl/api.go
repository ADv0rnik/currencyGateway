package impl

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func ReturnRouter() *Router {
	router := &Router{gin.Default()}
	return router
}

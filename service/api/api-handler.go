package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/enroll/", rt.wrap(rt.enrollNewUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

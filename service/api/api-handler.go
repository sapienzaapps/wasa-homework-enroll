package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/enroll/", rt.wrap(rt.enrollNewUser))

	rt.router.GET("/results/", rt.wrap(rt.listResults))
	rt.router.GET("/results/:studentid/git", rt.wrap(rt.getGitLog))
	rt.router.GET("/results/:studentid/openapi", rt.wrap(rt.getOpenAPILog))
	rt.router.GET("/results/:studentid/golang", rt.wrap(rt.getGoLog))
	rt.router.GET("/results/:studentid/vue", rt.wrap(rt.getVueLog))
	rt.router.GET("/results/:studentid/docker", rt.wrap(rt.getDockerLog))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

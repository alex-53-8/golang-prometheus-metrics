package rest

import (
	"context"
	"log"
	"net/http"

	gpmiddleware "github.com/carousell/gin-prometheus-middleware"
	"github.com/gin-gonic/gin"
)

type RestServer struct {
	restServer   *http.Server
	RouterGroups *RouterGroups
}

type RouterGroups struct {
	Public *gin.RouterGroup
}

func (s *RestServer) Start() {
	if err := s.restServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func (s *RestServer) Stop() {
	s.restServer.Shutdown(context.Background())
}

func CreateRestServer() *RestServer {
	var router = gin.Default()

	routerGroups := RouterGroups{
		Public: router.Group("/"),
	}

	p := gpmiddleware.NewPrometheus("gin")
	p.Use(router)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	return &RestServer{restServer: srv, RouterGroups: &routerGroups}
}

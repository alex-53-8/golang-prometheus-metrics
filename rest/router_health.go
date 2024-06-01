package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func AddHealthEndpoints(groups *RouterGroups) {
	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "ra_health_check",
		Help: "The total number of processed events",
	})

	groups.Public.GET("/health", func(c *gin.Context) {
		opsProcessed.Inc()
		c.JSON(http.StatusOK, gin.H{
			"status": "up and running",
		})
	})
}

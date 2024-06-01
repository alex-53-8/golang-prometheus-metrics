package processing

import (
	"rest_app_metrics/rest"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func AddEndpoints(groups *rest.RouterGroups) {

	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "ra_processing_endpoint",
		Help: "The total number of invocations",
	})

	groups.Public.GET("/processing", func(c *gin.Context) {
		opsProcessed.Inc()

		StartProcessing()
	})

}

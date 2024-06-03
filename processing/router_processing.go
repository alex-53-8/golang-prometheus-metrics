package processing

import (
	"net/http"

	"rest_app_metrics/rest"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ProcessingResponse struct {
	Id uuid.UUID `json:"id"`
}

func AddEndpoints(groups *rest.RouterGroups) {
	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "ra_processing_endpoint_count",
		Help: "The total number of invocations /processing endpoint",
	})

	groups.Public.PUT("/processing/create", func(c *gin.Context) {
		opsProcessed.Inc()

		id := StartProcessing()

		c.JSON(http.StatusOK, ProcessingResponse{Id: id})
	})

}

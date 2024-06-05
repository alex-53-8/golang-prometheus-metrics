package processing

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var processingDurationHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:    "ra_processing_time_total",
	Help:    "Total processing time histogram",
	Buckets: []float64{10, 20, 40, 80, 160, 320, 640, 1280, 3000, 5000, 7500, 10000, 15000},
})

var successfulProcessing = promauto.NewCounter(prometheus.CounterOpts{
	Name:        "ra_processing_count",
	Help:        "The total number of processed",
	ConstLabels: map[string]string{"status": "success"},
})
var failedProcessing = promauto.NewCounter(prometheus.CounterOpts{
	Name:        "ra_processing_count",
	Help:        "The total number of processed",
	ConstLabels: map[string]string{"status": "failed"},
})

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func StartProcessing() uuid.UUID {
	id, _ := uuid.NewRandom()

	go Process()

	return id
}

func recordProcessingTime(startTime time.Time) {
	time := float64(time.Now().UnixMilli()) - float64(startTime.UnixMilli())
	processingDurationHistogram.Observe(time)
}

func recordSuccessfullProcessing() {
	successfulProcessing.Inc()
}

func recordFailedProcessing() {
	failedProcessing.Inc()
}

func Process() {
	defer recordProcessingTime(time.Now())

	var sleepTimeSec = randRange(10, 16000)
	var sleepTime = time.Duration(sleepTimeSec) * time.Millisecond

	fmt.Println("processing a request")

	time.Sleep(sleepTime)

	if sleepTime%3 == 0 {
		fmt.Println("processing failed")
		recordFailedProcessing()
	} else {
		fmt.Println("processing succeed")
		recordSuccessfullProcessing()
	}

}

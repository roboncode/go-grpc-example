package prometheus

import (
	"example/internal/grpc"
	"example/util/log"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	//// Create a customized counter metric.
	//customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
	//	Name: "demo_server_say_hello_method_handle_count",
	//	Help: "Total number of RPCs handled on the server.",
	//}, []string{"name"})

	//docker run -p 9090:9090 prom/prometheus
	//docker run -p 9090:9090 -v /tmp/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
)

func NewServer(grpcServer grpc.Server) {
	// Create a HTTP server for prometheus.
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9092)}

	// Initialize all metrics.
	grpcMetrics.InitializeMetrics(grpcServer.Instance())

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Warnln("unable to start prometheus http server")
		}
	}()
}

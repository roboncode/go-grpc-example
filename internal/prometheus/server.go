package prometheus

import (
	"aaa/internal/grpc"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
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
)

func NewServer(grpcServer *grpc.Server) {
	// Create a HTTP server for prometheus.
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9092)}

	// Initialize all metrics.
	grpcMetrics.InitializeMetrics(grpcServer.Server())

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()
}

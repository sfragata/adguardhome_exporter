package collector

import (
	"log"

	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

// AdguardCollector struct that hold metrics collector insterface
type AdguardCollector struct {
	Collectors    []MetricCollector
	AdguardServer server.AdguardServer
}

// MetricCollector base for the metrics
type MetricCollector interface {
	collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error)
	describe() *prometheus.Desc
	metricType() prometheus.ValueType
}

// NewAdguardCollector constructor
func NewAdguardCollector(AdguardServer server.AdguardServer, adguardExporterVersion string) *AdguardCollector {
	return &AdguardCollector{

		Collectors: []MetricCollector{StatsMetrics{}, TopQueriesMetrics{}, TopBlockedMetrics{}, AdguardBuildInfoMetrics{
			version: adguardExporterVersion,
		}, FilteringStatusMetrics{}, TopClientsMetrics{}, DNSQueryTypesMetrics{}},
		AdguardServer: AdguardServer,
	}
}

// Describe method from prometheus Collector interface
func (pc *AdguardCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range pc.Collectors {
		ch <- metric.describe()
	}
}

// Collect method from prometheus Collector interface
func (pc *AdguardCollector) Collect(ch chan<- prometheus.Metric) {
	for _, collector := range pc.Collectors {

		if metrics, err := collector.collect(pc.AdguardServer); err == nil {
			for _, metric := range metrics {
				ch <- metric
			}
		} else {
			log.Print(err)
		}
	}
}

package collector

import (
	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

// topQueriedDomains represents the structure of the JSON
type topQueriedDomains struct {
	TopQueriedDomains []map[string]int `json:"top_queried_domains"`
}

// TopQueriesMetrics TopQueriesMetrics
type TopQueriesMetrics struct {
}

func (sm TopQueriesMetrics) describe() *prometheus.Desc {
	return prometheus.NewDesc("adguard_top_queried_domains", "show adguard top queried domains", []string{"domain"}, nil)
}
func (sm TopQueriesMetrics) metricType() prometheus.ValueType {
	return prometheus.GaugeValue
}
func (sm TopQueriesMetrics) collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error) {

	var queries topQueriedDomains

	if err := AdguardServer.SendRequest(statsEndpoint, &queries); err != nil {
		return nil, err
	}

	var metrics []prometheus.Metric

	for _, query := range queries.TopQueriedDomains {
		for key, value := range query {
			metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(value), key))
		}

	}

	return metrics, nil
}

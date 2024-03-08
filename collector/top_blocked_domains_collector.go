package collector

import (
	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

type topBlockedDomains struct {
	TopBlockedDomains []map[string]int `json:"top_blocked_domains"`
}

// TopBlockedMetrics TopQueriesMetrics
type TopBlockedMetrics struct {
}

func (sm TopBlockedMetrics) describe() *prometheus.Desc {
	return prometheus.NewDesc("adguard_top_blocked_domains", "show adguard top blocked domains", []string{"domain"}, nil)
}
func (sm TopBlockedMetrics) metricType() prometheus.ValueType {
	return prometheus.GaugeValue
}
func (sm TopBlockedMetrics) collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error) {

	var blockedDomains topBlockedDomains

	if err := AdguardServer.SendRequest(statsEndpoint, &blockedDomains); err != nil {
		return nil, err
	}

	var metrics []prometheus.Metric

	for _, query := range blockedDomains.TopBlockedDomains {
		for key, value := range query {
			metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(value), key))
		}

	}

	return metrics, nil
}

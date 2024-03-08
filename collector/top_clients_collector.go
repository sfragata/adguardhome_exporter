package collector

import (
	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

// topClients represents the structure of the JSON
type topClients struct {
	TopClients []map[string]int `json:"top_clients"`
}

// TopClientsMetrics TopClientsMetrics
type TopClientsMetrics struct {
}

func (sm TopClientsMetrics) describe() *prometheus.Desc {
	return prometheus.NewDesc("adguard_top_clients", "show adguard top clients", []string{"hostname"}, nil)
}
func (sm TopClientsMetrics) metricType() prometheus.ValueType {
	return prometheus.GaugeValue
}
func (sm TopClientsMetrics) collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error) {

	var clients topClients

	if err := AdguardServer.SendRequest(statsEndpoint, &clients); err != nil {
		return nil, err
	}

	var metrics []prometheus.Metric

	for _, query := range clients.TopClients {
		for key, value := range query {
			metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(value), key))
		}

	}

	return metrics, nil
}

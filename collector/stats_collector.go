package collector

import (
	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

const statsEndpoint = "control/stats"

// stats represents the structure of the JSON
type stats struct {
	NumDNSQueries           int     `json:"num_dns_queries"`
	NumBlockedFiltering     int     `json:"num_blocked_filtering"`
	NumReplacedSafebrowsing int     `json:"num_replaced_safebrowsing"`
	NumReplacedSafesearch   int     `json:"num_replaced_safesearch"`
	NumReplacedParental     int     `json:"num_replaced_parental"`
	AvgProcessingTime       float64 `json:"avg_processing_time"`
}

// StatsMetrics StatsMetrics
type StatsMetrics struct {
}

func (sm StatsMetrics) describe() *prometheus.Desc {
	return prometheus.NewDesc("adguard_stats", "show adguard stats", []string{"name"}, nil)
}
func (sm StatsMetrics) metricType() prometheus.ValueType {
	return prometheus.GaugeValue
}
func (sm StatsMetrics) collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error) {

	var stats stats

	if err := AdguardServer.SendRequest(statsEndpoint, &stats); err != nil {
		return nil, err
	}

	var metrics []prometheus.Metric

	metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(stats.NumBlockedFiltering), "num_blocked_filtering"))
	metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(stats.NumDNSQueries), "num_dns_queries"))
	metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(stats.NumReplacedParental), "num_replaced_parental"))
	metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(stats.NumReplacedSafebrowsing), "num_replaced_safebrowsing"))
	metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(stats.NumReplacedSafesearch), "num_replaced_safesearch"))
	metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(stats.AvgProcessingTime), "avg_processing_time"))

	return metrics, nil
}

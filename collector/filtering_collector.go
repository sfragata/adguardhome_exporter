package collector

import (
	"fmt"

	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

const filteringEndpoint = "control/filtering/status"

type fileringStatus struct {
	Filters []struct {
		URL         string `json:"url"`
		Name        string `json:"name"`
		LastUpdated string `json:"last_updated"`
		RulesCount  int    `json:"rules_count"`
		Enabled     bool   `json:"enabled"`
	} `json:"filters"`
	// WhitelistFilters interface{}   `json:"whitelist_filters"`
	// UserRules        []interface{} `json:"user_rules"`
	// Interval         int           `json:"interval"`
	// Enabled          bool          `json:"enabled"`
}

// FilteringStatusMetrics FilteringStatusMetrics
type FilteringStatusMetrics struct {
}

func (sm FilteringStatusMetrics) describe() *prometheus.Desc {
	return prometheus.NewDesc("adguard_filtering_status", "show adguard filters", []string{"url", "name", "last_update", "enable"}, nil)
}
func (sm FilteringStatusMetrics) metricType() prometheus.ValueType {
	return prometheus.GaugeValue
}
func (sm FilteringStatusMetrics) collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error) {

	var filteringStruct fileringStatus

	if err := AdguardServer.SendRequest(filteringEndpoint, &filteringStruct); err != nil {
		return nil, err
	}

	var metrics []prometheus.Metric

	for _, filters := range filteringStruct.Filters {
		var filterEnabled int = 0
		if filters.Enabled {
			filterEnabled = 1
		}

		metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(filters.RulesCount), filters.URL, filters.Name, filters.LastUpdated, fmt.Sprint(filterEnabled)))
	}

	return metrics, nil
}

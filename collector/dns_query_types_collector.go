package collector

import (
	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

const queryLogEndpoint = "control/querylog?limit=2000&response_status=all"

// dnsAnswer struct from LogData
type dnsAnswer struct {
	TTL   float64     `json:"ttl"`
	Type  string      `json:"type"`
	Value interface{} `json:"strubg"`
}

// logData struct, sub struct of LogStats to collect the dns stats from the log
type logData struct {
	Answer []dnsAnswer `json:"answer"`
}

// logStats struct for the Adguard log statistics JSON API corresponding model.
type logStats struct {
	Data   []logData `json:"data"`
	Oldest string    `json:"oldest"`
}

// DNSQueryTypesMetrics DNSQueryTypesMetrics
type DNSQueryTypesMetrics struct {
}

func (sm DNSQueryTypesMetrics) describe() *prometheus.Desc {
	return prometheus.NewDesc("adguard_dns_query_types", "show dns query types", []string{"type"}, nil)
}
func (sm DNSQueryTypesMetrics) metricType() prometheus.ValueType {
	return prometheus.GaugeValue
}
func (sm DNSQueryTypesMetrics) collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error) {

	var dnsStatsResponse logStats

	if err := AdguardServer.SendRequest(queryLogEndpoint, &dnsStatsResponse); err != nil {
		return nil, err
	}

	var metrics []prometheus.Metric
	// grouping and counting dns types

	dnsStats := make(map[string]int)

	for _, data := range dnsStatsResponse.Data {
		for _, answer := range data.Answer {
			dnsStats[answer.Type] += 1
		}

	}
	for key, value := range dnsStats {
		metrics = append(metrics, prometheus.MustNewConstMetric(sm.describe(), sm.metricType(), float64(value), key))
	}
	return metrics, nil
}

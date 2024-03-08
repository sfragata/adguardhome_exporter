package collector

import (
	"strconv"

	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus"
)

const statusEndpoint = "control/status"

type responseBuildInfo struct {
	Version string `json:"version"`
	// Language                   string   `json:"language"`
	// DNSAddresses               []string `json:"dns_addresses"`
	// DNSPort                    int      `json:"dns_port"`
	// HTTPPort                   int      `json:"http_port"`
	// ProtectionDisabledDuration int      `json:"protection_disabled_duration"`
	ProtectionEnabled bool `json:"protection_enabled"`
	// DHCPAvailable              bool     `json:"dhcp_available"`
	Running bool `json:"running"`
}

// AdguardBuildInfoMetrics AdguardBuildInfoMetrics
type AdguardBuildInfoMetrics struct {
	version string
}

func (adgim AdguardBuildInfoMetrics) describe() *prometheus.Desc {
	return prometheus.NewDesc("adguard_exporter_build_info", "A metric with a constant '1' value labeled by adguard version and adguardhome_exporter version from which adguard/adguard_exporter was built.", []string{"adguard_version", "adguard_exporter_version", "protection_enabled", "running"}, nil)
}
func (adgim AdguardBuildInfoMetrics) metricType() prometheus.ValueType {
	return prometheus.GaugeValue
}
func (adgim AdguardBuildInfoMetrics) collect(AdguardServer server.AdguardServer) ([]prometheus.Metric, error) {
	var responseBuildInfo responseBuildInfo
	if err := AdguardServer.SendRequest(statusEndpoint, &responseBuildInfo); err != nil {
		return nil, err
	}
	var protected int = 0
	if responseBuildInfo.ProtectionEnabled {
		protected = 1
	}

	var running int = 0
	if responseBuildInfo.Running {
		running = 1
	}

	metric := prometheus.MustNewConstMetric(adgim.describe(), adgim.metricType(), float64(1), responseBuildInfo.Version, adgim.version, strconv.Itoa(protected), strconv.Itoa(running))

	return []prometheus.Metric{metric}, nil
}

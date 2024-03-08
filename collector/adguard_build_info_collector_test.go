package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

const validAdguardBuildInfoJSON = "{\"version\": \"v0.107.44\",\"language\": \"\",\"dns_addresses\": [\"127.0.0.1\",\"::1\",\"192.168.1.44\",\"fe80::555f:c25a:c875:bfath0\"],\"dns_port\": 53,\"http_port\": 80,\"protection_disabled_duration\": 0,\"protection_enabled\": true,\"dhcp_available\": true,\"running\": true}"

func TestAdguardBuildInfoCollectorCollect(test *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validAdguardBuildInfoJSON)
	}))
	defer server.Close()
	adguardExporterVersion := "1.0"
	collector := NewAdguardCollector(newAdguardServer(server), "")
	expected := fmt.Sprintf(`
	# HELP adguard_exporter_build_info A metric with a constant '1' value labeled by adguard version and adguardhome_exporter version from which adguard/adguard_exporter was built.
	# TYPE adguard_exporter_build_info gauge
	adguard_exporter_build_info{adguard_exporter_version="%s",adguard_version="v0.107.44", protection_enabled="1", running="1"} 1
	`, adguardExporterVersion)
	collector.Collectors = []MetricCollector{AdguardBuildInfoMetrics{version: adguardExporterVersion}}

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		test.Errorf("Error:\n%s", err)
	}

}

func newAdguardServer(serverTest *httptest.Server) server.AdguardServer {

	hostPort := strings.Split(serverTest.Listener.Addr().String(), ":")

	port, _ := strconv.Atoi(hostPort[1])

	return server.AdguardServer{
		Host: hostPort[0],
		Port: port,
	}

}

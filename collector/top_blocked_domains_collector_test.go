package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestTopBlockedDomainsCollectorCollect(test *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validStatsJSON)
	}))
	defer server.Close()

	collector := NewAdguardCollector(newAdguardServer(server), "")
	expected := `
	# HELP adguard_top_blocked_domains show adguard top blocked domains
	# TYPE adguard_top_blocked_domains gauge
	adguard_top_blocked_domains{domain="www.google.com"} 10096
	adguard_top_blocked_domains{domain="logs.netflix.com"} 2307
	adguard_top_blocked_domains{domain="dit.whatsapp.net"} 791
	adguard_top_blocked_domains{domain="app-measurement.com"} 563
	`
	collector.Collectors = []MetricCollector{TopBlockedMetrics{}}

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		test.Errorf("Error:\n%s", err)
	}

}

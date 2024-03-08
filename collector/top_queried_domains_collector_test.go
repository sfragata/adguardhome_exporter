package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestTopQueriedDomainsCollectorCollect(test *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validStatsJSON)
	}))
	defer server.Close()

	collector := NewAdguardCollector(newAdguardServer(server), "")
	expected := `
	# HELP adguard_top_queried_domains show adguard top queried domains
	# TYPE adguard_top_queried_domains gauge
	adguard_top_queried_domains{domain="avs-alexa-na.amazon.com"} 28641
	adguard_top_queried_domains{domain="connectivity-check.ubuntu.com"} 3395
	adguard_top_queried_domains{domain="info.cspserver.net"} 3283
	adguard_top_queried_domains{domain="gateway.fe2.apple-dns.net"} 2524
	`
	collector.Collectors = []MetricCollector{TopQueriesMetrics{}}

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		test.Errorf("Error:\n%s", err)
	}

}

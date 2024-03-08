package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestTopClientsCollectorCollect(test *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validStatsJSON)
	}))
	defer server.Close()

	collector := NewAdguardCollector(newAdguardServer(server), "")
	expected := `
	# HELP adguard_top_clients show adguard top clients
	# TYPE adguard_top_clients gauge
	adguard_top_clients{hostname="192.168.1.1"} 137864
	`
	collector.Collectors = []MetricCollector{TopClientsMetrics{}}

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		test.Errorf("Error:\n%s", err)
	}

}

package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

const validStatsJSON = "{\"time_units\":\"days\",\"top_queried_domains\":[{\"avs-alexa-na.amazon.com\":28641},{\"connectivity-check.ubuntu.com\":3395},{\"info.cspserver.net\":3283},{\"gateway.fe2.apple-dns.net\":2524}],\"top_clients\":[{\"192.168.1.1\":137864}],\"top_blocked_domains\":[{\"www.google.com\":10096},{\"logs.netflix.com\":2307},{\"dit.whatsapp.net\":791},{\"app-measurement.com\":563}],\"top_upstreams_responses\":[{\"https://dns.google:443/dns-query\":75713},{\"192.168.1.1:53\":1325}],\"top_upstreams_avg_time\":[{\"192.168.1.1:53\":0.5649681954716981},{\"https://dns.google:443/dns-query\":0.022404251792954976}],\"dns_queries\":[0,0,0,7510,60544,59669,10141],\"blocked_filtering\":[0,0,0,427,3849,2471,666],\"replaced_safebrowsing\":[0,0,0,0,0,0,0],\"replaced_parental\":[0,0,0,0,0,0,0],\"num_dns_queries\":137864,\"num_blocked_filtering\":7413,\"num_replaced_safebrowsing\":0,\"num_replaced_safesearch\":10442,\"num_replaced_parental\":0,\"avg_processing_time\":0.018206}"

func TestStatsCollectorCollect(test *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validStatsJSON)
	}))
	defer server.Close()

	collector := NewAdguardCollector(newAdguardServer(server), "")
	expected := `
	# HELP adguard_stats show adguard stats
	# TYPE adguard_stats gauge
	adguard_stats{name="num_blocked_filtering"} 7413
	adguard_stats{name="num_dns_queries"} 137864
	adguard_stats{name="num_replaced_parental"} 0
	adguard_stats{name="num_replaced_safebrowsing"} 0
	adguard_stats{name="num_replaced_safesearch"} 10442
	adguard_stats{name="avg_processing_time"} 0.018206
	`
	collector.Collectors = []MetricCollector{StatsMetrics{}}

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		test.Errorf("Error:\n%s", err)
	}

}

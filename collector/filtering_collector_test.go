package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

const validFilteringJSON = "{\"filters\":[{\"url\":\"https://adguardteam.github.io/HostlistsRegistry/assets/filter_1.txt\",\"name\":\"AdGuard DNS filter\",\"last_updated\":\"2024-03-06T20:10:09-05:00\",\"id\":1,\"rules_count\":62772,\"enabled\":true},{\"url\":\"https://adguardteam.github.io/HostlistsRegistry/assets/filter_2.txt\",\"name\":\"AdAway Default Blocklist\",\"last_updated\":\"2024-03-06T20:10:09-05:00\",\"id\":2,\"rules_count\":6540,\"enabled\":true}],\"whitelist_filters\":null,\"user_rules\":[],\"interval\":24,\"enabled\":true}"

func TestFilteringStatusCollectorCollect(test *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validFilteringJSON)
	}))
	defer server.Close()

	collector := NewAdguardCollector(newAdguardServer(server), "")
	expected := `
	# HELP adguard_filtering_status show adguard filters
	# TYPE adguard_filtering_status gauge
	adguard_filtering_status{enable="1",last_update="2024-03-06T20:10:09-05:00",name="AdAway Default Blocklist",url="https://adguardteam.github.io/HostlistsRegistry/assets/filter_2.txt"} 6540
	adguard_filtering_status{enable="1",last_update="2024-03-06T20:10:09-05:00",name="AdGuard DNS filter",url="https://adguardteam.github.io/HostlistsRegistry/assets/filter_1.txt"} 62772
	`
	collector.Collectors = []MetricCollector{FilteringStatusMetrics{}}

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		test.Errorf("Error:\n%s", err)
	}

}

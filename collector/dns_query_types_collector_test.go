package collector

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

const validQueryLogJSON = "{\"data\":[{\"answer\":[{\"type\":\"CNAME\",\"value\":\"fcfd-prodp-outer.trafficmanager.net.\",\"ttl\":2519},{\"type\":\"A\",\"value\":\"52.252.110.74\",\"ttl\":300}],\"answer_dnssec\":false,\"cached\":false,\"client\":\"192.168.1.1\",\"client_info\":{\"whois\":{},\"name\":\"\",\"disallowed_rule\":\"192.168.1.1\",\"disallowed\":false},\"client_proto\":\"\",\"elapsedMs\":\"58.95299\",\"question\":{\"class\":\"IN\",\"name\":\"clientfd.family.microsoft.com\",\"type\":\"A\"},\"reason\":\"NotFilteredNotFound\",\"rules\":[],\"status\":\"NOERROR\",\"time\":\"2024-03-07T23:00:25.804844083-05:00\",\"upstream\":\"https://dns.google:443/dns-query\"},{\"answer_dnssec\":false,\"cached\":false,\"client\":\"192.168.1.1\",\"client_info\":{\"whois\":{},\"name\":\"\",\"disallowed_rule\":\"192.168.1.1\",\"disallowed\":false},\"client_proto\":\"\",\"elapsedMs\":\"523.613333\",\"question\":{\"class\":\"IN\",\"name\":\"lb._dns-sd._udp.0.1.168.192.in-addr.arpa\",\"type\":\"PTR\"},\"reason\":\"NotFilteredNotFound\",\"rules\":[],\"status\":\"NXDOMAIN\",\"time\":\"2024-03-07T23:00:21.617205104-05:00\",\"upstream\":\"192.168.1.1:53\"}],\"oldest\":\"2024-03-07T23:00:21.617205104-05:00\"}"

func TestDNSQueryTypesCollectorCollect(test *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, validQueryLogJSON)
	}))
	defer server.Close()

	collector := NewAdguardCollector(newAdguardServer(server), "")
	expected := `
	# HELP adguard_dns_query_types show dns query types
	# TYPE adguard_dns_query_types gauge
	adguard_dns_query_types{type="CNAME"} 1
	adguard_dns_query_types{type="A"} 1
	`
	collector.Collectors = []MetricCollector{DNSQueryTypesMetrics{}}

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		test.Errorf("Error:\n%s", err)
	}

}

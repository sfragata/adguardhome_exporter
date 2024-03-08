package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com.sfragata/adguardhome_exporter/collector"
	"github.com.sfragata/adguardhome_exporter/server"
	"github.com/integrii/flaggy"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// These variables will be replaced by real values when do gorelease
var (
	version = "none"
	date    string
	commit  string
)

func main() {

	info := fmt.Sprintf(
		"%s\nDate: %s\nCommit: %s\nOS: %s\nArch: %s",
		version,
		date,
		commit,
		runtime.GOOS,
		runtime.GOARCH,
	)

	flaggy.SetName("adguardhome_exporter")
	flaggy.SetDescription("Prometheus exporter for Adguard home")
	flaggy.SetVersion(info)

	var adguardHost = "127.0.0.1"
	flaggy.String(&adguardHost, "H", "host", "Adguard home address")

	var adguardPort = 80
	flaggy.Int(&adguardPort, "p", "port", "Adguard home port")

	var adguardToken = os.Getenv("ADGUARD_HOME_TOKEN") // username:password in base64
	flaggy.String(&adguardToken, "t", "token", "Adguard home token (if ADGUARD_HOME_TOKEN env variable is set, don't need to pass it)")

	var metricsPort = "9311"
	flaggy.String(&metricsPort, "l", "listen-address", "Adguard home exporter metrics port")

	flaggy.Parse()

	client := &http.Client{
		Timeout: server.HTTPTimeout * time.Second,
	}

	AdguardServer := server.AdguardServer{
		Host:       adguardHost,
		Port:       adguardPort,
		Token:      adguardToken,
		HTTPClient: *client,
	}

	err := prometheus.Register(collector.NewAdguardCollector(AdguardServer, version))
	if err != nil {
		log.Fatalf("Can't register collectors: %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("starting adguardhome_export [:%s]", metricsPort)
	err = http.ListenAndServe(":"+metricsPort, nil)
	if err != nil {
		log.Fatalf("Can't start server %s", metricsPort)
	}

}

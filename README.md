# adguardhome_exporter
[Adguard home](https://github.com/AdguardTeam/AdGuardHome) prometheus exporter

![Golang CI](https://github.com/sfragata/adguardhome_exporter/workflows/Golang%20CI/badge.svg)

## Installation

### Mac

```
brew tap sfragata/tap

brew install sfragata/tap/adguardhome_exporter
```

### Linux and Windows

get latest release [here](https://github.com/sfragata/adguardhome_exporter/releases)

## Usage

```
adguardhome_exporter - Prometheus exporter for Adguard home

  Flags: 
       --version          Displays the program version string.
    -h --help             Displays help with available flag, subcommand, and positional value parameters.
    -H --host             Adguard home address (default: 127.0.0.1)
    -p --port             Adguard home port (default: 80)
    -t --token            Adguard home token (if ADGUARD_HOME_TOKEN env variable is set, don't need to pass it)
    -l --listen-address   Adguard home exporter metrics port (default: 9311)
```    
## Output
```
# HELP adguard_dns_query_types show dns query types
# TYPE adguard_dns_query_types gauge
adguard_dns_query_types{type="A"} 3182
adguard_dns_query_types{type="AAAA"} 356
adguard_dns_query_types{type="CNAME"} 2533
# HELP adguard_exporter_build_info A metric with a constant '1' value labeled by adguard version and adguardhome_exporter version from which adguard/adguard_exporter was built.
# TYPE adguard_exporter_build_info gauge
adguard_exporter_build_info{adguard_exporter_version="v1.0.0",adguard_version="v0.107.45",protection_enabled="1",running="1"} 1
# HELP adguard_filtering_status show adguard filters
# TYPE adguard_filtering_status gauge
adguard_filtering_status{enable="1",last_update="2024-03-07T21:01:58-05:00",name="AdAway Default Blocklist",url="https://adguardteam.github.io/HostlistsRegistry/assets/filter_2.txt"} 6540
adguard_filtering_status{enable="1",last_update="2024-03-07T21:01:58-05:00",name="AdGuard DNS filter",url="https://adguardteam.github.io/HostlistsRegistry/assets/filter_1.txt"} 62819
# HELP adguard_stats show adguard stats
# TYPE adguard_stats gauge
adguard_stats{name="avg_processing_time"} 0.023511
adguard_stats{name="num_blocked_filtering"} 16765
adguard_stats{name="num_dns_queries"} 340190
adguard_stats{name="num_replaced_parental"} 0
adguard_stats{name="num_replaced_safebrowsing"} 0
adguard_stats{name="num_replaced_safesearch"} 25312
# HELP adguard_top_blocked_domains show adguard top blocked domains
# TYPE adguard_top_blocked_domains gauge
adguard_top_blocked_domains{domain="4421.api.swrve.com"} 38
adguard_top_blocked_domains{domain="youtubei.googleapis.com"} 333
# HELP adguard_top_clients show adguard top clients
# TYPE adguard_top_clients gauge
adguard_top_clients{hostname="192.168.2.10"} 340190
# HELP adguard_top_queried_domains show adguard top queried domains
# TYPE adguard_top_queried_domains gauge
adguard_top_queried_domains{domain="10.195.225.13.in-addr.arpa"} 2296
adguard_top_queried_domains{domain="111.195.225.13.in-addr.arpa"} 2296
```
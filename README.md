# plex_exporter
[Plex](https://www.plex.tv) prometheus exporter

![Golang CI](https://github.com/sfragata/plex_exporter/workflows/Golang%20CI/badge.svg)

## Installation

### Mac

```
brew tap sfragata/tap

brew install sfragata/tap/plex_exporter
```

### Linux and Windows

get latest release [here](https://github.com/sfragata/plex_exporter/releases)

## Usage

```
plex_exporter - Prometheus exporter for plex

  Flags:
       --version          Displays the program version string.
    -h --help             Displays help with available flag, subcommand, and positional value parameters.
    -H --host             Plex address (default: 127.0.0.1)
    -p --port             Plex port (default: 32400)
    -t --token            Plex token (if PLEX_TOKEN env variable is set, it will be used)
    -l --listen-address   Plex exporter metrics port (default: 2112)
```    
## Output
```
# HELP plex_active_sessions_count show number active sessions
# TYPE plex_active_sessions_count gauge
plex_active_sessions_count 2
# HELP plex_active_sessions_count_user_device show number active sessions by user and device
# TYPE plex_active_sessions_count_user_device gauge
plex_active_sessions_count_user_device{device="OSX",user="xxx"} 1
plex_active_sessions_count_user_device{device="iPhone",user="yyy"} 1
# HELP plex_library_count show number medias
# TYPE plex_library_count gauge
plex_library_count{name="Rock",type="artist"} 6
plex_library_count{name="Movies",type="movie"} 97
plex_library_count{name="Series",type="show"} 2
```
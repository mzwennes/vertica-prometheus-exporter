# Vertica Prometheus Exporter

Exports important `Vertica` metrics in Prometheus format. 

## Usage
```
go run main.go -h              
Usage of vertica-prometheus-exporter:
  -database string
        Vertica database name (default "vertica")
  -host string
        Vertica hostname (default "localhost")
  -listen string
        Address to listen on (default "0.0.0.0:8080")
  -location string
        Metrics path (default "/metrics")
  -password string
        Vertica password (default "dbadmin")
  -port int
        Vertica port (default 5433)
  -user string
        Vertica username (default "dbadmin")
```

## Build


## Development
A local Vertica database can be deployed for testing using the following command:
```
docker run --name vertica --network host -p 5433:5433 -d jbfavre/vertica:9.2.0-7_debian-8
```
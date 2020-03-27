# Vertica Prometheus Exporter

Exports important `Vertica` metrics in Prometheus format. 

## Usage
```
go run main.go -h              
Usage of vertica-prometheus-exporter:
  -location string
        Metrics path (default "/metrics")
  -listen string
        Address to listen on (default "0.0.0.0:8080")
  -db_name string
        Vertica database name (default "vertica")
  -db_host string
        Vertica hostname (default "localhost")
  -db_password string
        Vertica password (default "dbadmin")
  -db_port int
        Vertica port (default 5433)
  -db_user string
        Vertica username (default "dbadmin")
```

Running this application will serve Prometheus metrics in a fashion such as below:
```
vertica_node_state{node_id="45035996273704978", node_name="v_docker_node0001"} 1
vertica_wos_used_bytes 0
vertica_last_good_epoch 16
vertica_node_count 1
vertica_catalog_revision_number 726
vertica_current_epoch 17
vertica_ahm_epoch 16
vertica_current_fault_tolerance 0
vertica_wos_row_count 0
vertica_refresh_epoch -1
vertica_ros_row_count 0
vertica_total_used_bytes 0
vertica_designed_fault_tolerance 0
vertica_node_down_count 0
vertica_ros_used_bytes 0
vertica_total_row_count 0
```

## Build
The application can be built locally with:

```
go build
```

Or you can create a Docker image:

```
docker build -t vertica-prometheus-exporter .
```

The newly created Docker image can then be started (note the use of `ENV variables`):

```
docker run --rm --network host --name vertica-prometheus-exporter \
      -e DB_NAME=docker \
      vertica-prometheus-exporter:latest
```

## Development

A local Vertica database can be deployed for testing using the following command:

```
docker run --name vertica --network host -p 5433:5433 -d jbfavre/vertica:9.2.0-7_debian-8
```

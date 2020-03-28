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
vertica_request_duration_ms{user_name="dbadmin"} 16275
vertica_memory_acquired_mb{user_name="dbadmin"} 18284
vertica_request_duration_ms{user_name="testuser"} 216
vertica_memory_acquired_mb{user_name="testuser"} 500
vertica_pool_running_query_count{node_name="v_docker_node0001", pool_name="general"} 0
vertica_pool_memory_inuse_kb{node_name="v_docker_node0001", pool_name="general"} 0
vertica_pool_memory_borrowed_kb{node_name="v_docker_node0001", pool_name="general"} 0
vertica_pool_memory_inuse_kb{node_name="v_docker_node0001", pool_name="sysquery"} 0
```

## Build locally
The application can be built locally with:

```
go build
```

Build the Docker image:

```
docker build -t vertica-prometheus-exporter .
```

## Running with Docker

Either use the locally built image, or download a pre-made image from the Github container repository.

```
docker pull docker.pkg.github.com/zwennesm/vertica-prometheus-exporter/vertica-prometheus-exporter:latest
```

Then run the docker image with optional `environment variables`:

```
docker run --rm --network host --name vertica-prometheus-exporter \
      -e DB_NAME=docker \
      vertica-prometheus-exporter:latest
```

## Running the binaries
You can either run the Docker image or download and start the application directly. Each new release contains binaries for different architectures.

Example of downloading and running the latest version (`linux_386`):
```
wget https://github.com/zwennesm/vertica-prometheus-exporter/releases/download/v0.3/vertica-prometheus-exporter_v0.3_linux_386.tar.gz

tar -xvf vertica-prometheus-exporter_v0.3_linux_386.tar.gz

./vertica-prometheus-exporter
```

## Development

A local Vertica database can be deployed for testing using the following command:

```
docker run --name vertica --network host -p 5433:5433 -d jbfavre/vertica:9.2.0-7_debian-8
```

Note: The default database name is `docker`.
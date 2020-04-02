package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
	_ "github.com/vertica/vertica-sql-go"

	"github.com/jmoiron/sqlx"
	"github.com/zwennesm/vertica-prometheus-exporter/monitoring"
)

func main() {
	location := flag.String("location", "/metrics", "Metrics path")
	listen := flag.String("listen", "0.0.0.0:8080", "Address to listen on")
	dbUsername := flag.String("db_user", "dbadmin", "Vertica username")
	dbPassword := flag.String("db_password", "dbadmin", "Vertica password")
	dbHost := flag.String("db_host", "localhost", "Vertica hostname")
	dbPort := flag.Int("db_port", 5433, "Vertica port")
	dbName := flag.String("db_name", "vertica", "Vertica database name")
	flag.Parse()

	connString := fmt.Sprintf("vertica://%v:%v@%v:%v/%v", *dbUsername, *dbPassword, *dbHost, *dbPort, *dbName)
	db, err := sqlx.Connect("vertica", connString)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	serveMetrics(*location, *listen, *db)
}

// Serve Vertica metrics at chosen address and url.
func serveMetrics(location, listen string, db sqlx.DB) {

	h := func(w http.ResponseWriter, r *http.Request) {
		metrics := monitoring.NewPrometheusMetrics(db)
		for _, obj := range metrics {
			metric := obj.ToMetric()
			for key, value := range metric {
				fmt.Fprintf(w, "%s %d\n", key, value)
			}
		}
	}

	http.HandleFunc(location, h)
	log.Printf("starting serving metrics at %s%s", listen, location)
	err := http.ListenAndServe(listen, nil)
	log.Fatal(err)
}

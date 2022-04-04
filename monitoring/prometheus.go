package monitoring

import (
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

// PrometheusMetric maps a struct to a Prometheus valid map.
type PrometheusMetric interface {
	ToMetric() map[string]float32
}

func NewPrometheusMetrics(db sqlx.DB) []PrometheusMetric {
	var metrics []PrometheusMetric

	for _, state := range NewNodeState(&db) {
		metrics = append(metrics, state)
	}
	for _, rejection := range NewPoolRejections(&db) {
		metrics = append(metrics, rejection)
	}
	for _, queryRequest := range NewQueryRequests(&db) {
		metrics = append(metrics, queryRequest)
	}
	for _, usage := range NewPoolUsage(&db) {
		metrics = append(metrics, usage)
	}
	for _, schemaSize := range NewSchemaSize(&db) {
		metrics = append(metrics, schemaSize)
	}
	for _, licenseSize := range NewLicenseSize(&db) {
		metrics = append(metrics, licenseSize)
	}
	metrics = append(metrics, NewVerticaSystem(&db))

	return metrics
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts all string values to snake case.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

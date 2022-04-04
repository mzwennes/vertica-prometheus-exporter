package monitoring

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// QueryRequest lists query performance metrics on the username level.
type QueryRequest struct {
	UserName          string `db:"user_name"`
	RequestDurationMS int    `db:"request_duration_ms"`
	MemoryAcquiredMB  int    `db:"memory_acquired_mb"`
}

// NewQueryRequests returns query performance for all users.
func NewQueryRequests(db *sqlx.DB) []QueryRequest {
	sql := `
	SELECT
		user_name, 
		SUM(COALESCE(request_duration_ms,0))::INT request_duration_ms, 
		SUM(COALESCE(memory_acquired_mb,0))::INT memory_acquired_mb 
	FROM v_monitor.query_requests 
	GROUP BY user_name;
	`

	queryRequests := []QueryRequest{}
	err := db.Select(&queryRequests, sql)
	if err != nil {
		log.Fatal(err)
	}

	return queryRequests
}

// ToMetric converts QueryRequest to a Map.
func (qr QueryRequest) ToMetric() map[string]float32 {
	metrics := map[string]float32{}

	username := fmt.Sprintf("user_name=%q", qr.UserName)
	metrics[fmt.Sprintf("vertica_request_duration_ms{%s}", username)] = float32(qr.RequestDurationMS)
	metrics[fmt.Sprintf("vertica_memory_acquired_mb{%s}", username)] = float32(qr.MemoryAcquiredMB)

	return metrics
}

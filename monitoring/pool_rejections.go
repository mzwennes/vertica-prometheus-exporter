package monitoring

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// PoolRejection shows the amount of resource pool rejections per node.
type PoolRejection struct {
	NodeName       string `db:"node_name"`
	PoolName       string `db:"pool_name"`
	RejectionCount int    `db:"rejection_count"`
}

// NewPoolRejections returns a list of resource pool rejections from Vertica.
func NewPoolRejections(db *sqlx.DB) []PoolRejection {
	sql := `
	SELECT 
		node_name, 
		pool_name,
		rejection_count
	FROM v_monitor.resource_rejections`

	rejections := []PoolRejection{}
	err := db.Select(&rejections, sql)
	if err != nil {
		log.Fatal(err)
	}

	return rejections
}

// ToMetric converts PoolRejection to a Map.
func (pr PoolRejection) ToMetric() map[string]int {
	metrics := map[string]int{}

	node := fmt.Sprintf("node_name=%q", pr.NodeName)
	pool := fmt.Sprintf("pool_name=%q", pr.PoolName)
	metrics[fmt.Sprintf("vertica_pool_rejection_count{%s, %s}", node, pool)] = pr.RejectionCount

	return metrics
}

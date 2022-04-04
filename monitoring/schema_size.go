package monitoring

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// PoolUsage shows gneral resource pool usage stats.
type SchemaSize struct {
	SchemaName   string `db:"schema_name"`
	SchemaSizeGB string `db:"schema_size_gb"`
}

// NewPoolUsage returns a list of pool usage stats.
func NewSchemaSize(db *sqlx.DB) []SchemaSize {
	sql := `
	SELECT 
		anchor_table_schema as schema_name,
		SUM(used_bytes) / 1024/1024/1024 AS schema_size_gb
	FROM v_monitor.projection_storage
	GROUP BY anchor_table_schema;
	`

	usage := []SchemaSize{}
	err := db.Select(&usage, sql)
	if err != nil {
		log.Fatal(err)
	}

	return usage
}

// ToMetric converts PoolUsage to a Map.
func (usage SchemaSize) ToMetric() map[string]float32 {
	metrics := map[string]float32{}

	schemaSize, err := strconv.ParseFloat(usage.SchemaSizeGB, 32)
	if err != nil {
		fmt.Println(err)
	}

	schemaName := fmt.Sprintf("schema=%q", usage.SchemaName)
	metrics[fmt.Sprintf("vertica_schema_size{%s}", schemaName)] = float32(schemaSize)

	return metrics
}

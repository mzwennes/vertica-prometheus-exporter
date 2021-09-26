package monitoring

import (
	"fmt"
	"log"

	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
)

// VerticaSystem shows important system values such as the epoch and fault tolerance levels.
type VerticaSystem struct {
	CurrentEpoch           int `db:"current_epoch"`
	AhmEpoch               int `db:"ahm_epoch"`
	LastGoodEpoch          int `db:"last_good_epoch"`
	RefreshEpoch           int `db:"refresh_epoch"`
	DesignedFaultTolerance int `db:"designed_fault_tolerance"`
	NodeCount              int `db:"node_count"`
	NodeDownCount          int `db:"node_down_count"`
	CurrentFaultTolerance  int `db:"current_fault_tolerance"`
	CatalogRevisionNumber  int `db:"catalog_revision_number"`
	WosUsedBytes           int `db:"wos_used_bytes"`
	WosRowCount            int `db:"wos_row_count"`
	RosUsedBytes           int `db:"ros_used_bytes"`
	RosRowCount            int `db:"ros_row_count"`
	TotalUsedBytes         int `db:"total_used_bytes"`
	TotalRowCount          int `db:"total_row_count"`
}

// NewVerticaSystem returns a new instance of VerticaSystem
func NewVerticaSystem(db *sqlx.DB) VerticaSystem {
	sql := `SELECT * FROM system`

	system := VerticaSystem{}
	err := db.Get(&system, sql)
	if err != nil {
		log.Fatal(err)
	}

	return system
}

// ToMetric converts VerticaSystem to a Map.
func (sys VerticaSystem) ToMetric() map[string]float32 {
	metrics := map[string]float32{}

	for k, v := range structs.Map(sys) {
		metrics[fmt.Sprintf("vertica_%s", ToSnakeCase(k))] = float32(v.(int))
	}

	return metrics
}

package monitoring

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// NodeState contains information about each node in a Vertica cluster.
type NodeState struct {
	NodeID    string `db:"node_id""`
	NodeName  string `db:"node_name"`
	NodeState int    `db:"node_state"`
}

// NewNodeState returns the status for each node in the Vertica cluster.
func NewNodeState(db *sqlx.DB) []NodeState {
	sql := `
	SELECT 
		node_id, 
		node_name, 
		(node_state='UP')::INT node_state 
	FROM v_catalog.nodes`

	nodeState := []NodeState{}
	err := db.Select(&nodeState, sql)
	if err != nil {
		log.Fatal(err)
	}

	return nodeState
}

// ToMetric converts NodeState to a Map.
func (ns NodeState) ToMetric() map[string]int {
	metrics := map[string]int{}

	id := fmt.Sprintf("node_id=%q", ns.NodeID)
	name := fmt.Sprintf("node_name=%q", ns.NodeName)
	metrics[fmt.Sprintf("vertica_node_state{%s, %s}", id, name)] = ns.NodeState

	return metrics
}

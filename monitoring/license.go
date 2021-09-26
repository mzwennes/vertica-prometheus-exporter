package monitoring

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// PoolUsage shows gneral resource pool usage stats.
type LicenseSize struct {
	AuditLicenseSize string `db:"audit_license_size"`
}

// NewPoolUsage returns a list of pool usage stats.
func NewLicenseSize(db *sqlx.DB) []LicenseSize {
	sql := `
	SELECT audit_license_size();
	`

	licenseSize := []LicenseSize{}
	err := db.Select(&licenseSize, sql)
	if err != nil {
		log.Fatal(err)
	}

	return licenseSize
}

// ToMetric converts PoolUsage to a Map.
func (licenseSize LicenseSize) ToMetric() map[string]float32 {
	metrics := map[string]float32{}

	rows := strings.Split(licenseSize.AuditLicenseSize, "\n")
	rawDataSizeStr := strings.Replace(strings.Fields(rows[0])[3], "TB", "", -1)
	rawDataSize, err := strconv.ParseFloat(rawDataSizeStr, 32)
	if err != nil {
		fmt.Println(err)
	}

	licenseSizeStr := strings.Replace(strings.TrimSpace(strings.Split(rows[1], ":")[1]), "TB", "", -1)
	licenseSizeRow, err := strconv.ParseFloat(licenseSizeStr, 32)
	if err != nil {
		fmt.Println(err)
	}

	utilizationStr := strings.Replace(strings.TrimSpace(strings.Split(rows[2], ":")[1]), "%", "", -1)
	utilization, err := strconv.ParseFloat(utilizationStr, 32)
	if err != nil {
		fmt.Println(err)
	}

	metrics["vertica_license_raw_data_size_tb"] = float32(rawDataSize)
	metrics["vertica_license_size_tb"] = float32(licenseSizeRow)
	metrics["vertica_license_utilization"] = float32(utilization)

	return metrics
}

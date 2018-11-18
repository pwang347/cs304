package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// QueryNumberOfOrganizationsForService returns count of organizations using a service
func QueryNumberOfOrganizationsForService(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response    = SQLResponse{}
		tx          *sql.Tx
		serviceName string
	)
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT COUNT(DISTINCT organizationName) as organizationCount FROM ServiceSubscriptionTransaction "+
		"WHERE serviceName = ? GROUP BY serviceName, organizationName;", serviceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryNumberOfInstancesPerRegion returns count of instances per region
func QueryNumberOfInstancesPerRegion(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response    = SQLResponse{}
		tx          *sql.Tx
		serviceName string
	)
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT regionName, COUNT(*) as count FROM ServiceInstance "+
		"WHERE serviceName = ? GROUP BY regionName;", serviceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryNumberOfVirtualMachinesPerRegion returns count of instances per region
func QueryNumberOfVirtualMachinesPerRegion(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response    = SQLResponse{}
		tx          *sql.Tx
		serviceName string
	)
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT COUNT(*) as count FROM VirtualMachine "+
		"WHERE virtualMachineServiceName = ? GROUP BY regionName;", serviceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryWeeklyPurchasesForService returns the count of purchases for service this week
func QueryWeeklyPurchasesForService(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response    = SQLResponse{}
		tx          *sql.Tx
		serviceName string
	)
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT COUNT(*) as purchases FROM ServiceSubscriptionTransaction "+
		"WHERE serviceName = ? AND processedTimestamp > (NOW() - INTERVAL 7 DAY);", serviceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

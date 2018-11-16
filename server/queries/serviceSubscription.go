package queries

import (
	"database/sql"
	"encoding/json"
	"github.com/pwang347/cs304/server/common"
	"net/url"
	"strconv"
)

// TODO: When subscribed to a service, we want to create a corresponding transaction.. Might want to put
// service subscription table and transaction tables together? So when a subscription is made, we update both
// transaction and service subscription.

// CreateServiceSubscription creates a new service subscription
func CreateServiceSubscription(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		serviceTypeStr   string
		serviceType		 int
		serviceName		 string
		description		 string
		organizationName string
		activeUntilStr	 string
		activeUntil		 int
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if serviceTypeStr, err = common.GetRequiredParam(params, "serviceType"); err != nil {
		return
	}
	if serviceType, err = strconv.Atoi(serviceTypeStr); err != nil {
		return
	}
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if description, err = common.GetRequiredParam(params, "description"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if activeUntilStr, err = common.GetRequiredParam(params, "activeUntil"); err != nil {
		return
	}
	if activeUntil, err = strconv.Atoi(activeUntilStr); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO ServiceSubscription (type,serviceName,description, organizationName, activeUntil) VALUES(?,?,?,?,?);",
		serviceType, serviceName, description, organizationName, activeUntil); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// DeleteServiceSubscription deletes a service subscription
func DeleteServiceSubscription(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		serviceTypeStr   string
		serviceType		 int
		serviceName		 string
	)

	if tx, err = db.Begin(); err != nil {
		return
	}
	if serviceTypeStr, err = common.GetRequiredParam(params, "serviceType"); err != nil {
		return
	}
	if serviceType, err = strconv.Atoi(serviceTypeStr); err != nil {
		return
	}
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM ServiceSubscription WHERE type=? AND serviceName=?;",
		serviceType, serviceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	if response.AffectedRows, err = result.RowsAffected(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}
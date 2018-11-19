package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"

	"github.com/pwang347/cs304/server/common"
)

// QueryAllServices returns all service rows
func QueryAllServices(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{}
		tx       *sql.Tx
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM Service;"); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// Get all service types given a service name
func GetServiceTypes(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{}
		tx       *sql.Tx
		name     string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx,
		"SELECT * FROM Service S, ServiceServiceTypePairs SSTP, ServiceType ST "+
			"WHERE S.name = ? AND S.name = SSTP.serviceName AND SSTP.serviceType = ST.type;", name); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// Gets all subscriptions/transactions for one service within the given organization
func GetServiceSubscriptions(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		serviceName      string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if serviceName, err = common.GetRequiredParam(params, "serviceName"); err != nil {
		return
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx,
		"SELECT * FROM Service, ServiceSubscriptionTransaction WHERE name = serviceName "+
			"AND organizationName = ? AND serviceName = ?;", organizationName, serviceName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// Complex division query that gets all services that have all types
func GetAllServicesWithAllTypes(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{}
		tx       *sql.Tx
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx,
		"select distinct s.name, s.description, s.isPreview, s.isEnabled, s.isVirtualMachineService, s.imageUrl "+
			"from Service s, ServiceServiceTypePairs sst1 "+
			"where s.name = sst1.serviceName and not exists ("+
			"select st.type from ServiceType st where st.type not in ("+
			"select sst2.serviceType from ServiceServiceTypePairs sst2 where sst1.serviceName = sst2.serviceName));"); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

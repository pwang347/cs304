package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pwang347/cs304/server/common"
)

// CreateOrganization creates a new organization
func CreateOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result              sql.Result
		response            = SQLResponse{}
		tx                  *sql.Tx
		name                string
		contactEmailAddress string
		tmpRowsAffected     int64
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if contactEmailAddress, err = common.GetRequiredParam(params, "contactEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO Organization (name,createdTimestamp,contactEmailAddress) VALUES(?,NOW(),?);",
		name, contactEmailAddress); err != nil {
		tx.Rollback()
		return
	}
	if result, err = tx.Exec("INSERT INTO UserOrganizationPairs (organizationName,userEmailAddress) VALUES(?,?);",
		name, contactEmailAddress); err != nil {
		return
	}
	if tmpRowsAffected, err = result.RowsAffected(); err != nil {
		return
	}
	response.AffectedRows += tmpRowsAffected
	if err = tx.Commit(); err != nil {
		return
	}
	if tmpRowsAffected, err = result.RowsAffected(); err != nil {
		return
	}
	response.AffectedRows += tmpRowsAffected
	data, err = json.Marshal(response)
	return
}

// DeleteOrganization deletes an organization
func DeleteOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result   sql.Result
		response = SQLResponse{}
		tx       *sql.Tx
		name     string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if name, err = common.GetRequiredParam(params, "name"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM Organization WHERE name=?;", name); err != nil {
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

// AddUserToOrganization adds a user to an existing organization
func AddUserToOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		userEmailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("INSERT INTO UserOrganizationPairs (organizationName,userEmailAddress) VALUES(?,?);", organizationName, userEmailAddress); err != nil {
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

// RemoveUserFromOrganization removes a user from an existing organization
func RemoveUserFromOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		userEmailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if result, err = tx.Exec("DELETE FROM UserOrganizationPairs WHERE organizationName=? AND userEmailAddress=?;",
		organizationName, userEmailAddress); err != nil {
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

// QueryUserOrganizations queries all organizations a user belongs to
func QueryUserOrganizations(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		userEmailAddress string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM UserOrganizationPairs UOP, Organization O "+
		"WHERE UOP.organizationName = O.name AND userEmailAddress = ?;", userEmailAddress); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryOrganizationUsersNotInGroup queries all users an organization contains not in an access group
func QueryOrganizationUsersNotInGroup(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		accessGroupName  string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if accessGroupName, err = common.GetRequiredParam(params, "accessGroupName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM UserOrganizationPairs UOP, User U "+
		"WHERE UOP.userEmailAddress = U.emailAddress AND UOP.organizationName = ? "+
		"AND NOT EXISTS (select * from UserAccessGroupPairs UAGP where UAGP.userEmailAddress = U.emailAddress AND UAGP.accessGroupName = ?);",
		organizationName, accessGroupName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryOrganizationUsers queries all users an organization contains
func QueryOrganizationUsers(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM UserOrganizationPairs UOP, User U "+
		"WHERE UOP.userEmailAddress = U.emailAddress AND UOP.organizationName = ? ",
		organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// QueryUserNotInOrganization queries all users not in an organization
func QueryUserNotInOrganization(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if response.Data, response.AffectedRows, err = common.QueryJSON(tx, "SELECT * FROM User U "+
		"WHERE NOT EXISTS (select * from UserOrganizationPairs UOP where UOP.userEmailAddress = U.emailAddress AND UOP.organizationName = ?);",
		organizationName); err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
	data, err = json.Marshal(response)
	return
}

// UpdateUserOrganizationPairs updates a user organization pair
func UpdateUserOrganizationPairs(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		result           sql.Result
		response         = SQLResponse{}
		tx               *sql.Tx
		organizationName string
		userEmailAddress string
		isAdminStr       string
		isAdmin          bool
	)

	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	if organizationName, err = common.GetRequiredParam(params, "organizationName"); err != nil {
		return
	}
	if userEmailAddress, err = common.GetRequiredParam(params, "userEmailAddress"); err != nil {
		return
	}
	if isAdminStr, err = common.GetRequiredParam(params, "isAdmin"); err != nil {
		return
	}
	if isAdmin, err = strconv.ParseBool(isAdminStr); err != nil {
		return
	}
	if result, err = tx.Exec("UPDATE UserOrganizationPairs SET isAdmin = ? WHERE organizationName=? AND userEmailAddress=?;",
		isAdmin, organizationName, userEmailAddress); err != nil {
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

package queries

import (
	"database/sql"
	"encoding/json"
	"net/url"
)

// CreateServiceInstance creates a new service instance
func CreateServiceInstance(db *sql.DB, params url.Values) (data []byte, err error) {
	var (
		response = SQLResponse{Rows: 0}
	)

	// TODO

	data, err = json.Marshal(response)
	return
}

package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

const (

	// DefaultMySQLUser represents the default user
	DefaultMySQLUser = "root"

	// DefaultMySQLPassword represents the default password
	DefaultMySQLPassword = ""

	// DefaultMySQLHost represents the default host
	DefaultMySQLHost = "localhost"

	// DefaultMySQLPort represents the default port
	DefaultMySQLPort = 3306

	// DefaultMySQLDbName represents the default database name
	DefaultMySQLDbName = "cs304"
)

// QueryJSON performs a query operation and returns the results as JSON
func QueryJSON(tx *sql.Tx, sqlString string, args ...interface{}) (jsonResult string, numRows int64, err error) {
	var (
		rows *sql.Rows
	)
	rows, err = tx.Query(sqlString, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		numRows++
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return
	}
	jsonResult = string(jsonData)
	fmt.Println(jsonResult)
	return
}

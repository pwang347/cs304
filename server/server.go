package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pwang347/cs304/server/common"
	"github.com/pwang347/cs304/server/queries"
)

const (
	// DefaultServerPort represents the default port number for server
	DefaultServerPort = 8080

	// ErrorNotFound represents the error for URL not found
	ErrorNotFound = "not-found"
)

var (
	db                  *sql.DB
	organizationQueries = map[string]query{
		"createUser": queries.CreateUser,
	}
	serviceQueries = map[string]query{
		"createServiceInstance": queries.CreateServiceInstance,
	}
	billingQueries = map[string]query{
		"createTransaction": queries.CreateTransaction,
	}
)

type (
	query         func(db *sql.DB, params url.Values) ([]byte, error)
	errorResponse struct {
		Error string `json:"error"`
	}
)

func mapJSONEndpoints(queries map[string]query) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		vars := mux.Vars(r)
		if cmd, ok := queries[vars["query"]]; ok {
			var (
				data []byte
				err  error
			)
			w.Header().Set("Content-Type", "application/json")
			data, err = cmd(db, r.URL.Query())
			if err != nil {
				if err.Error() == ErrorNotFound {
					http.NotFound(w, r)
					return
				}
				data, _ = json.Marshal(errorResponse{Error: err.Error()})
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(data)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		}
		http.NotFound(w, r)
	}
}

func main() {
	var (
		port          = flag.Int("port", DefaultServerPort, "webserver port")
		mysqlUser     = flag.String("user", common.DefaultMySQLUser, "mysql user")
		mysqlPassword = flag.String("password", common.DefaultMySQLPassword, "mysql password")
		mysqlHost     = flag.String("mysqlHost", common.DefaultMySQLHost, "mysql host")
		mysqlPort     = flag.Int("mysqlPort", common.DefaultMySQLPort, "mysql port")
		mysqlDbName   = flag.String("dbName", common.DefaultMySQLDbName, "mysql database name")
		err           error
	)

	flag.Parse()

	if len(*mysqlPassword) > 0 {
		*mysqlPassword = ":" + *mysqlPassword
	}

	mySQLConnectionString := fmt.Sprintf(
		"%s%s@tcp(%s:%d)/%s",
		*mysqlUser,
		*mysqlPassword,
		*mysqlHost,
		*mysqlPort,
		*mysqlDbName)

	if db, err = sql.Open("mysql", mySQLConnectionString); err != nil {
		panic(err)
	}

	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/organization/{query}", mapJSONEndpoints(organizationQueries))
	r.HandleFunc("/api/service/{query}", mapJSONEndpoints(serviceQueries))
	r.HandleFunc("/api/billing/{query}", mapJSONEndpoints(billingQueries))
	http.Handle("/", r)

	fmt.Println(fmt.Sprintf("Starting webserver at http://0.0.0.0:%d...", *port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.LoggingHandler(os.Stdout, r)))
}

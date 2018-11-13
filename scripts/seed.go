package main

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pwang347/cs304/server/common"
)

var (
	counter int = 0
	db      *sql.DB
)

func generateGroupName() interface{} {
	// TODO
	return "Test group name"
}

func generateOrganizationName() interface{} {
	// TODO
	return "Test organization name"
}

func generateOsName() interface{} {
	// TODO
	return "Test OS"
}

func generateVersion() interface{} {
	// TODO
	return "Test version"
}

func generateCreditCardNumber() interface{} {
	// TODO
	return "12345678901234567890"
}

func generateCvc() interface{} {
	// TODO
	return "123"
}

func generateTimestamp() interface{} {
	// TODO
	return time.Now()
}

func generateCreditCardType() interface{} {
	// TODO
	return "Test card type"
}

func generateEventLogNumber() interface{} {
	// TODO
	return 0
}

func generateEventLogData() interface{} {
	// TODO
	return "Test event log data"
}

func generateEmailAddress() interface{} {
	// TODO
	counter++
	return fmt.Sprintf("test%d@email.com", counter)
}

func generateFirstName() interface{} {
	// TODO
	return "TestFirst"
}

func generateLastName() interface{} {
	// TODO
	return "TestLast"
}

func generatePasswordHash() interface{} {
	// TODO
	return "TestHash"
}

func generateBoolean() interface{} {
	// TODO
	return true
}

func generatePhoneNumber() interface{} {
	// TODO
	counter++
	return fmt.Sprintf("%dISATEST", counter)
}

type attributeGenerator struct {
	name      string
	generator func() interface{}
}

const (
	maxGenerations int = 100
	minGenerations int = 50
)

func main() {
	var (
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

	rand.Seed(time.Now().UTC().UnixNano())

	// TODO: generate fake data for each table here
	if err = generate("User",
		attributeGenerator{name: "emailAddress", generator: generateEmailAddress},
		attributeGenerator{name: "firstName", generator: generateFirstName},
		attributeGenerator{name: "lastName", generator: generateLastName},
		attributeGenerator{name: "passwordHash", generator: generatePasswordHash},
		attributeGenerator{name: "isAdmin", generator: generateBoolean},
		attributeGenerator{name: "twoFactorPhoneNumber", generator: generatePhoneNumber}); err != nil {
		panic(err)
	}

	fmt.Println("Finished generating fake data.")
}

func generate(table string, generators ...attributeGenerator) error {
	ntimes := rand.Intn(maxGenerations-minGenerations) + minGenerations
	return generateNTimes(table, ntimes, generators...)
}

func generateNTimes(table string, ntimes int, generators ...attributeGenerator) (err error) {
	fmt.Println(fmt.Sprintf("Generating %d rows for %s...", ntimes, table))
	var tx *sql.Tx
	substitutions := strings.TrimSuffix(strings.Repeat("?,", len(generators)), ",")
	for i := 0; i < ntimes; i++ {
		if tx, err = db.Begin(); err != nil {
			return err
		}
		if _, err = tx.Exec(fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s);", table, strings.Join(extractNames(generators), ","), substitutions),
			extractValues(generators)...); err != nil {
			tx.Rollback()
			return err
		}
		err = tx.Commit()
	}
	fmt.Println("Done")
	return err
}

func extractNames(generators []attributeGenerator) []string {
	names := make([]string, len(generators))
	for i := 0; i < len(generators); i++ {
		names[i] = generators[i].name
	}
	return names
}

func extractValues(generators []attributeGenerator) []interface{} {
	values := make([]interface{}, len(generators))
	for i := 0; i < len(generators); i++ {
		values[i] = generators[i].generator()
	}
	return values
}

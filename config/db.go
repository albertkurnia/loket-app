package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"  
)

const (

	// maxIdleConn - maximum idle connection
	maxIdleConn = 10

	// maxOpenConn = maximum open connection
	maxOpenConn = 10

	// thresholdAttempt - maximum attempt trying to connect to database
	thresholdAttempt = 10
)

var (

	// DBMasterType - database postgre master type
	DBMasterType DBType = 1

	// DBSlaveType - database postgre slave type
	DBSlaveType DBType = 2
)

// DBConnector - abstraction to connect database
type DBConnector interface {
	DataSource(dbType DBType) string
}

// DBType - type that will be described as types of database that system provides
// eg: Master and slave
type DBType int

// PSQLInfo - PSQL configuration information
type PSQLInfo struct{}

// DataSource - get data source from configuration
func (pq *PSQLInfo) DataSource(dbType DBType) string {

	// Get connection based on database type
	// 1 -> Master
	// 2 -> Slave
	// default -> Master
	switch dbType {

	case 1:

		// Return datasource for database master
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable fallback_application_name=%s",
			os.Getenv("WRITE_DB_HOST"), os.Getenv("WRITE_DB_PORT"), os.Getenv("WRITE_DB_USER"),
			os.Getenv("WRITE_DB_PASSWORD"), os.Getenv("WRITE_DB_NAME"), os.Getenv("WRITE_DB_FALLBACK"))
		return psqlInfo
	case 2:

		// Return datasource for database slave
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable fallback_application_name=%s",
			os.Getenv("READ_DB_HOST"), os.Getenv("READ_DB_PORT"), os.Getenv("READ_DB_NAME"), os.Getenv("READ_DB_USER"),
			os.Getenv("READ_DB_PASSWORD"), os.Getenv("READ_DB_FALLBACK"))
	default:

		// Return datasource for database master
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable fallback_application_name=%s",
			os.Getenv("WRITE_DB_HOST"), os.Getenv("WRITE_DB_PORT"), os.Getenv("WRITE_DB_NAME"), os.Getenv("WRITE_DB_USER"),
			os.Getenv("WRITE_DB_PASSWORD"), os.Getenv("WRITE_DB_FALLBACK"))
	}
}

// CreateDBConn - create database connection
func CreateDBConn(connector DBConnector, dbType DBType) (*sql.DB, error) {

	// Trying to connect to postgreSQL server
	ds := connector.DataSource(dbType)
	db, err := sql.Open("postgres", ds)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	// Attempt to connect if thre's an error when first time connection
	if err != nil {
		for i := 1; i <= thresholdAttempt; i++ {
			err = db.Ping()
			if err == nil {
				break
			}

			// Set sleep for a moment to make interval connection
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	// After reached maximum threshold of attemption,
	// make a panic process, because the system can not establish connection
	// to database server
	if err != nil {
		return nil, err
	}

	// Note: maximum open connection must be greater or equals than maximum open connection
	// by default, if max idle connection is greter than max open connection, idle connection
	// will use max open connection value.

	// Set maximum idle connection
	db.SetMaxIdleConns(maxIdleConn)

	// Set maximum idle connection
	db.SetMaxOpenConns(maxOpenConn)

	return db, nil
}

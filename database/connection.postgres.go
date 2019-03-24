package database

import (
	"distributorBE/constants"
	"fmt"

	"github.com/jinzhu/gorm"
)

var (

	// DbConn Connection
	DbConn    *gorm.DB
	dbUser    string
	dbPass    string
	dbName    string
	dbHost    string
	dbPort    string
	dbType    string
	sslMode   string
	debugMode bool
)

func init() {

	dbType = "POSTGRES"
	dbUser = constants.DbUser
	dbPass = constants.DbPass
	dbName = constants.DbName
	dbHost = constants.DbHost
	dbPort = constants.DbPort
	sslMode = "disable"
	debugMode = true
	fmt.Println("Starting connec to Database ...")
	if DbOpen() != nil {
		fmt.Println("Failed connected to Database !")
	} else {
		fmt.Println("Database connected !")
	}

}

// DbOpen Open connection
func DbOpen() error {

	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, sslMode)

	DbConn, errDb := gorm.Open("postgres", args)

	if errDb != nil {
		fmt.Println("Open db error ", errDb)
		return errDb
	}

	if errPing := DbConn.DB().Ping(); errPing != nil {
		return errPing
	} else {
		fmt.Println("ping to Database success ! ")
	}

	return nil
}

// GetDbConnection ...
func GetDbConnection() *gorm.DB {

	if errPing := DbConn.DB().Ping(); errPing != nil {
		fmt.Println("Error ping to Database server")
	} else {
		fmt.Println("Database still alive ! ")
	}
	DbConn.LogMode(debugMode)
	return DbConn

}

/**
DB Instance Singleton - Access to data models persistence

Uses gorm.io/gorm, and currently supports:
- MySQL
... and that's it ...
Other supported DBs:
- PostgreSQL
- SQLite
- SQL Server
- Clickhouse
*/
package db

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type DBClient struct {
	db *gorm.DB
}

func (client *DBClient) AutoMigrate() {
	client.db.AutoMigrate(&Topic{})
	client.db.AutoMigrate(&Subscriber{})
}

var dbInstance *DBClient

func GetInstance(dbType string, url string, database string, username string, password string) *DBClient {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			fmt.Printf("Creating DB Instance Now %s, %s, %s \n", dbType, url, database)

			// Check on https://gorm.io/docs/connecting_to_the_database.html how to support additional DB Engines
			var err error
			var db *gorm.DB
			switch strings.ToUpper(dbType) {
			case "MYSQL":
				db, err = NewMYSQLDB(url, database, username, password)
				if err != nil {
					log.Fatalf("Failed to connect to DB %s", err)
				}
				dbInstance = &DBClient{db}
			default:
				log.Fatalf("Database type %s not supported", dbType)
			}

		} else {
			fmt.Println("DB Instance already created-1")
		}
	} else {
		fmt.Println("DB Instance already created-2")
	}
	return dbInstance
}

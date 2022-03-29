package service

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresClient *gorm.DB
var postgresClientSingleton sync.Once

func CreatePostgresClient(host, username, password string, port uint, dbName string) {
	postgresClientSingleton.Do(func() {
		//dsn := fmt.Sprintf("host=%s user=%s  dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, username, "pass", "word", password, dbName, port)
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v, sslmode=disable TimeZone=Asia/Shanghai", host, username, password, dbName, port)
		var e error
		postgresClient, e = gorm.Open(postgres.Open(dsn))
		if e != nil {
			log.Default().Panic("Postgres DB initiation failed. Reason", e.Error())
		}
	})
}

func GetPostgresDBClient() *gorm.DB {
	return postgresClient
}

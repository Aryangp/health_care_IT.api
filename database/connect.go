package database

import (
	"fmt"
	"log"

	"github.com/Aryangp/goRest/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=1978 dbname=healthdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                          // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("database connected")
	}

	fmt.Println("connected to database")

	return db
}
func InitialMigration() {
	db := GetDatabase()
	db.AutoMigrate(model.User{})
}

// func Closedatabase(connection *gorm.DB) {
// 	sqldb := connection.DB()
// 	sqldb.Close()
// }

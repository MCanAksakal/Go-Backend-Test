package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"database/sql"

	_ "github.com/lib/pq"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB
var err error

func InitGormPostgres() {
	MPosGORM, err = gorm.Open("postgres", "user=xxx dbname=yyy password=zzz sslmode=disable")
	if err != nil {
		panic(err)
	}
}

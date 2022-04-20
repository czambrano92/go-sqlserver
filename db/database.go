package db

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var dsn = "sqlserver://gouser:cerveros.2022@localhost:1433?database=patentes"

var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("ERROR EN LA CONEXION", err)
		panic(err)
	} else {
		fmt.Println("conexion exitosa")
		return db
	}
}

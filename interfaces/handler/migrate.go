package handler

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    DB *gorm.DB
)

func init() {
    DB = CreateConnection()
}

func CreateConnection() *gorm.DB {
    DBMS     := "postgres"
    USER     := "root"
    PASS     := "pass"
    PROTOCOL := "tcp(192.168.33.50:5432)"
    DBNAME   := "test"
    CONNECT  := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?parseTime=True&loc=Local"
    
    db, err := gorm.Open(DBMS, CONNECT)

    if err != nil {
        panic(err)
    }
    fmt.Printf("[RYON-debug] Connected %s\n", DBMS)
    return db
}

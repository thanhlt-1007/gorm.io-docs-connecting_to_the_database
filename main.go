package main

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

const MYSQL_USER = "root"
const MYSQL_PASSWORD = "Aa@123456"
const MYSQL_HOST = "127.0.0.1:3306"
const MYSQL_DATABASE = "new_db_1"

var MYSQL_DSN_CONNECT_MYSQL = fmt.Sprintf("%s:%s@tcp(%s)/", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST)
var MYSQL_DSN_CONNECT_DATABASE = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DATABASE)

func createDatabaseIfNotExists() {
    fmt.Println("createDatabase")
    db, err := gorm.Open(mysql.Open(MYSQL_DSN_CONNECT_MYSQL), &gorm.Config{})
    if err != nil {
        fmt.Printf("createDatabase err: %v\n", err)
        panic(err)
    }

    sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", MYSQL_DATABASE)
    fmt.Printf("Exec sql: %s\n", sql)
    executedDb := db.Exec(sql)
    if executedDb.Error != nil {
        fmt.Printf("Exec Error: %v\n", executedDb.Error)
        panic(executedDb.Error)
    }

    fmt.Printf("Create database %s\n", MYSQL_DATABASE)
}

func main() {
    createDatabaseIfNotExists()

    _, err := gorm.Open(mysql.Open(MYSQL_DSN_CONNECT_DATABASE), &gorm.Config{})
    if err != nil {
        fmt.Printf("gorm.Open error: %v\n", err)
        panic(err)
    }

    fmt.Println("gorm.Open success")
}

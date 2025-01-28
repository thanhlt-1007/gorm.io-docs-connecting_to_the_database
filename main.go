package main

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-connecting_to_the_database/configs"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var MYSQL_DSN_CONNECT_MYSQL = fmt.Sprintf("%s:%s@tcp(%s)/",
    configs.MYSQL_USER, configs.MYSQL_PASSWORD, configs.MYSQL_HOST)

var MYSQL_DSN_CONNECT_DATABASE = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    configs.MYSQL_USER, configs.MYSQL_PASSWORD, configs.MYSQL_HOST, configs.MYSQL_DATABASE)

func createDatabaseIfNotExists() {
    fmt.Println("createDatabase")
    db, err := gorm.Open(mysql.Open(MYSQL_DSN_CONNECT_MYSQL), &gorm.Config{})
    if err != nil {
        fmt.Printf("createDatabase err: %v\n", err)
        panic(err)
    }

    sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", configs.MYSQL_DATABASE)
    fmt.Printf("Exec sql: %s\n", sql)
    executedDb := db.Exec(sql)
    if executedDb.Error != nil {
        fmt.Printf("Exec Error: %v\n", executedDb.Error)
        panic(executedDb.Error)
    }

    fmt.Printf("Create database %s\n", configs.MYSQL_DATABASE)
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

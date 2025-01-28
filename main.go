package main

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func main() {
    MYSQL_USER := "root"
    MYSQL_PASSWORD := "Aa@123456"
    MYSQL_HOST := "127.0.0.1:3306"
    MYSQL_DATABASE := "db"

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DATABASE)
    _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Printf("gorm.Open error: %v\n", err)
        panic(err)
    }

    fmt.Println("gorm.Open success")
}

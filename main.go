package main

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func main() {
    dsn := "root:Aa@123456@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
    _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Printf("gorm.Open error: %v\n", err)
        panic(err)
    }
}

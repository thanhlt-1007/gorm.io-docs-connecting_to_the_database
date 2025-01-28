package database

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-connecting_to_the_database/configs"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
    db := connectMySql()
    createDatabaseIfNotExists(db)

    db = openDatabase()

    return db
}

func connectMySql() *gorm.DB {
    fmt.Println("Connecting MySQL...")

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/",
        configs.MYSQL_USER,
        configs.MYSQL_PASSWORD,
        configs.MYSQL_HOST,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Printf("Connect MySQL error: %v\n", err)
        panic(err)
    }

    fmt.Println("Connect MySQL success")
    fmt.Println()
    return db
}

func createDatabaseIfNotExists(db *gorm.DB) *gorm.DB {
    fmt.Printf("Creating database %s..\n", configs.MYSQL_DATABASE)

    sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET %s COLLATE %s;",
        configs.MYSQL_DATABASE,
        configs.MYSQL_CHARACTER_SET,
        configs.MYSQL_COLLATE,
    )
    db = db.Exec(sql)

    err := db.Error
    if err != nil {
        fmt.Printf("Create database error: %v\n", err)
        panic(err)
    }

    fmt.Printf("Create database %s success\n\n",
        configs.MYSQL_DATABASE,
    )
    return db
}

func openDatabase() *gorm.DB {
    fmt.Printf("Opening database %s...\n",
        configs.MYSQL_DATABASE,
    )

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
        configs.MYSQL_USER,
        configs.MYSQL_PASSWORD,
        configs.MYSQL_HOST,
        configs.MYSQL_DATABASE,
        configs.MYSQL_CHARACTER_SET,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Printf("Open database error: %v\n", err)
        panic(err)
    }

    fmt.Printf("Open database %s success\n\n",
        configs.MYSQL_DATABASE,
    )
    return db
}

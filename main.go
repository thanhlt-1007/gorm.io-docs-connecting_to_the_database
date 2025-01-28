package main

import (
    "github.com/thanhlt-1007/gorm.io-docs-connecting_to_the_database/database"
)

func main() {
    db := database.InitDatabase()
    database.AutoMigrate(db)
}

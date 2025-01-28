package database

import (
    "gorm.io/gorm"
    "github.com/thanhlt-1007/gorm.io-docs-connecting_to_the_database/models"
)

func AutoMigrate(db *gorm.DB) {
    db.AutoMigrate(&models.User{})
}

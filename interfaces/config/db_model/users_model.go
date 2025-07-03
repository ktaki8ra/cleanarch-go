package db_model

import (
    "gorm.io/gorm"
)

type Users struct {
    gorm.Model
    ID uint `gorm:"primaryKey"`
    UserId string
    Email string
    EncryptedPassword string
}

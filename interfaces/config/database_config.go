package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type DBConfig struct {
    Host string
    User string
    DBName string
    Port int
}

func LoadDBConfig() DBConfig {
    return DBConfig{
        Host: "localhost",
        User: "kodai",
        DBName: "cleanarch_sample_go",
        Port: 5432,
    }
}

type DBConnection struct {
    DB *gorm.DB
}
func NewDBConnection(dbConfig DBConfig) (*DBConnection, error) {
    dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo", dbConfig.Host, dbConfig.User, dbConfig.DBName, dbConfig.Port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return &DBConnection{DB: db}, nil
}

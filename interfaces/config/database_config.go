package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os/user"
    "log"
    "os"
)

type DBConfig struct {
    Host string
    User string
    DBName string
    Port int
}

func getCurrentUser() (*user.User, error) {
    return user.Current()
}

func LoadDBConfig() DBConfig {
    currentUser, err := getCurrentUser()
    if err != nil {
        log.Fatalf("Failed to get current User info")
        os.Exit(1)
    }
    return DBConfig{
        Host: "localhost",
        User: currentUser.Username,
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

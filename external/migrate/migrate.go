package main

import (
    "os"
    "github.com/ktaki8ra/cleanarch-go/interfaces/config"
    "github.com/ktaki8ra/cleanarch-go/interfaces/config/db_model"
)

func main() {
    dbConfig := config.LoadDBConfig()
    dbConn, err := config.NewDBConnection(dbConfig)
    if err != nil {
        os.Exit(1)
    }
    dbConn.DB.AutoMigrate(&db_model.Users{})
}

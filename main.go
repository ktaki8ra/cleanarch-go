package main

import (
    "os"
    "github.com/ktaki8ra/cleanarch-go/interfaces/config"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/controller"
)

func main() {
    err := controller.New(config.LoadDI(config.LoadDBConfig())).Run(config.LoadHttpConfig())
    if err != nil {
        os.Exit(1)
    }
}

package config

import (
    "os"

    "github.com/ktaki8ra/cleanarch-go/usecases/service"
    "github.com/ktaki8ra/cleanarch-go/domain/domain_repository"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/gateway"
)

type DIModules struct {
    CryptoService service.CryptoService
    UserRepository domain_repository.UserRepository
}

func LoadDI(dbConfig DBConfig) DIModules {
    // DI for cryptoService
    cryptoServiceImpl := &gateway.CryptoServiceImpl{}

    // DI for userRepository
    dbConn, err := NewDBConnection(dbConfig)
    if err != nil {
        os.Exit(1)
    }
    userRepositoryImpl := &gateway.UserRepositoryImpl{DB: dbConn.DB}

    // Aggregate DI Modules
    return DIModules {
        CryptoService: cryptoServiceImpl,
        UserRepository: userRepositoryImpl,
    }
}

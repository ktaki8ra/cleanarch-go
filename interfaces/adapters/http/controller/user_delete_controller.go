package controller

import (
    "github.com/labstack/echo/v4"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/decoder"
    "github.com/ktaki8ra/cleanarch-go/interfaces/config"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/validator"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/presenter"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"

    "log"
)

func UserDeleteController(di config.DIModules) echo.HandlerFunc {
    return func(c echo.Context) error {

        userDeleteRequestJson, decodeJsonErr := decoder.DecodeUserDeleteRequestJson(c)
        if decodeJsonErr.Err != nil {
            log.Printf("User Delete Request Json couldn't be decoded.")
            decoderErrorResponse := presenter.DecoderJsonErrorConvert(decodeJsonErr)
            return c.JSON(decoderErrorResponse.StatusCode, decoderErrorResponse)
        }

        userDeleteInputData, userDeleteValidationErr := validator.ValidateUserDeleteData(userDeleteRequestJson)
        if userDeleteValidationErr.Err != nil {
            log.Printf("User Delete Request Data failed Validation.")
            validationErrorResponse := presenter.ValidationErrorConvert(userDeleteValidationErr)
            return c.JSON(validationErrorResponse.StatusCode, validationErrorResponse)
        }

        userDeleteUseCase := usecase.NewUserDeleteUseCase(
            di.CryptoService,
            di.UserRepository,
        )
        userDeleteOutputData, useCaseErr := userDeleteUseCase.Execute(userDeleteInputData)
        if useCaseErr.Err != nil {
            log.Printf("User Delete UseCase failed.")
            usecaseErrorResponse := presenter.UseCaseErrorConvert(useCaseErr)
            return c.JSON(usecaseErrorResponse.StatusCode, usecaseErrorResponse)
        }

        userDeleteResponseJson := presenter.UserDeleteSuccessConvert(userDeleteOutputData)
        return c.JSON(userDeleteResponseJson.StatusCode, userDeleteResponseJson)
    }
}

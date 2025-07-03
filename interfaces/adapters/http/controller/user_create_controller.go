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

func UserCreateController(di config.DIModules) echo.HandlerFunc {
    return func(c echo.Context) error {

        userCreateRequestJson, decodeJsonErr := decoder.DecodeUserCreateRequestJson(c)
        if decodeJsonErr.Err != nil {
            log.Printf("User Create Request Json couldn't be decoded.")
            decoderErrorResponse := presenter.DecoderJsonErrorConvert(decodeJsonErr)
            return c.JSON(decoderErrorResponse.StatusCode, decoderErrorResponse)
        }

        userCreateInputData, userCreateValidationErr := validator.ValidateUserCreateData(userCreateRequestJson)
        if userCreateValidationErr.Err != nil {
            log.Printf("User Create Request Data failed Validation.")
            validationErrorResponse := presenter.ValidationErrorConvert(userCreateValidationErr)
            return c.JSON(validationErrorResponse.StatusCode, validationErrorResponse)
        }

        userCreateUseCase := usecase.NewUserCreateUseCase(
            di.CryptoService,
            di.UserRepository,
        )
        userCreateOutputData, useCaseErr := userCreateUseCase.Execute(userCreateInputData)
        if useCaseErr.Err != nil {
            log.Printf("User Create UseCase failed.")
            usecaseErrorResponse := presenter.UseCaseErrorConvert(useCaseErr)
            return c.JSON(usecaseErrorResponse.StatusCode, usecaseErrorResponse)
        }

        userCreateResponseJson := presenter.UserCreateSuccessConvert(userCreateOutputData)
        return c.JSON(userCreateResponseJson.StatusCode, userCreateResponseJson)

    }
}

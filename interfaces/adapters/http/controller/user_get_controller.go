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

func UserGetController(di config.DIModules) echo.HandlerFunc {
    return func(c echo.Context) error {

        userGetRequestJson, decodeJsonErr := decoder.DecodeUserGetRequestJson(c)
        if decodeJsonErr.Err != nil {
            log.Printf("User Get Request Json couldn't be decoded.")
            decoderErrorResponse := presenter.DecoderJsonErrorConvert(decodeJsonErr)
            return c.JSON(decoderErrorResponse.StatusCode, decoderErrorResponse)
        }

        userGetInputData, userGetValidationErr := validator.ValidateUserGetData(userGetRequestJson)
        if userGetValidationErr.Err != nil {
            log.Printf("User Get Request Data failed Validation.")
            validationErrorResponse := presenter.ValidationErrorConvert(userGetValidationErr)
            return c.JSON(validationErrorResponse.StatusCode, validationErrorResponse)
        }

        userGetUseCase := usecase.NewUserGetUseCase(
            di.UserRepository,
        )
        userGetOutputData, useCaseErr := userGetUseCase.Execute(userGetInputData)
        if useCaseErr.Err != nil {
            log.Printf("User Get UseCase failed.")
            usecaseErrorResponse := presenter.UseCaseErrorConvert(useCaseErr)
            return c.JSON(usecaseErrorResponse.StatusCode, usecaseErrorResponse)
        }

        userGetResponseJson := presenter.UserGetSuccessConvert(userGetOutputData)
        return c.JSON(userGetResponseJson.StatusCode, userGetResponseJson)
    }
}

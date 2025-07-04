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

func UserUpdateController(di config.DIModules) echo.HandlerFunc {
    return func(c echo.Context) error {

        userUpdateRequestJson, decodeJsonErr := decoder.DecodeUserUpdateRequestJson(c)
        if decodeJsonErr.Err != nil {
            log.Printf("User Delete Request Json couldn't be decoded.")
            decoderErrorResponse := presenter.DecoderJsonErrorConvert(decodeJsonErr)
            return c.JSON(decoderErrorResponse.StatusCode, decoderErrorResponse)
        }

        userUpdateInputData, userUpdateValidationErr := validator.ValidateUserUpdateData(userUpdateRequestJson)
        if userUpdateValidationErr.Err != nil {
            log.Printf("User Delete Request Data failed Validation.")
            validationErrorResponse := presenter.ValidationErrorConvert(userUpdateValidationErr)
            return c.JSON(validationErrorResponse.StatusCode, validationErrorResponse)
        }

        userUpdateUseCase := usecase.NewUserUpdateUseCase(
            di.CryptoService,
            di.UserRepository,
        )
        userUpdateOutputData, useCaseErr := userUpdateUseCase.Execute(userUpdateInputData)
        if useCaseErr.Err != nil {
            log.Printf("User Delete UseCase failed.")
            usecaseErrorResponse := presenter.UseCaseErrorConvert(useCaseErr)
            return c.JSON(usecaseErrorResponse.StatusCode, usecaseErrorResponse)
        }

        userUpdateResponseJson := presenter.UserUpdateSuccessConvert(userUpdateOutputData)
        return c.JSON(userUpdateResponseJson.StatusCode, userUpdateResponseJson)
    }
}

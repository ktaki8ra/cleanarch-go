package presenter

import (
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/http/http_json"
    "github.com/ktaki8ra/cleanarch-go/interfaces/adapters/adapter_error"
    "github.com/ktaki8ra/cleanarch-go/usecases/usecase"
)

func DecoderJsonErrorConvert(decodeJsonError adapter_error.DecodeJsonError) http_json.ErrorResponseJson {
    return http_json.ErrorResponseJson {
        StatusCode: decodeJsonError.StatusCode,
        Msg: decodeJsonError.Msg,
    }
}

func ValidationErrorConvert(validationError adapter_error.ValidationError) http_json.ErrorResponseJson {
    return http_json.ErrorResponseJson {
        StatusCode: validationError.StatusCode,
        Msg: validationError.Msg,
    }
}

func UseCaseErrorConvert(usecaseError usecase.UseCaseError) http_json.ErrorResponseJson {
    return http_json.ErrorResponseJson {
        StatusCode: usecaseError.StatusCode,
        Msg: usecaseError.Msg,
    }
}

package adapter_error

import (
    "strconv"
)

type DecodeJsonError struct {
    StatusCode int
    Msg string
    Err error
}

type ValidationError struct {
    StatusCode int
    Msg string
    Err error
}

func (dje *DecodeJsonError) Error() string {
    return "<status_code>" + strconv.Itoa(dje.StatusCode) + "<Message>" + dje.Msg + " <Error Value>" + dje.Err.Error()
}
func (dje *DecodeJsonError) Unwrap() error {
    return dje.Err
}

func (ve *ValidationError) Error() string {
    return "<status_code>" + strconv.Itoa(ve.StatusCode) + "<Message>" + ve.Msg + " <Error Value>" + ve.Err.Error()
}
func (ve *ValidationError) Unwrap() error {
    return ve.Err
}

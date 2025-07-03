package usecase

import (
    "strconv"
)

type UseCaseError struct {
    StatusCode int
    Msg string
    Err error
}

func (uce *UseCaseError) Error() string {
    return "Status " + strconv.Itoa(uce.StatusCode) +" " + uce.Msg + " " + uce.Err.Error()
}

func (uce *UseCaseError) Unwrap() error {
    return uce.Err
}

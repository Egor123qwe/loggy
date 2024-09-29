package model

import "errors"

var (
	NotInitializedErr = errors.New("logger not initialized")

	AuthHeadersInvalidErr = errors.New("invalid auth headers")
	BadCredentialsErr     = errors.New("invalid credentials")

	BadRequestErr = errors.New("bad request")

	UnauthorizedErr = errors.New("unauthorized")
)

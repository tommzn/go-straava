package strava

import (
	"errors"
	"time"

	"golang.org/x/oauth2"
)

// NewTokenSourceMock returns a local token source for testing.
// If shouldReturnError is true Token() method will always return with an error.
func newTokenSourceMock(shouldReturnError bool) *tokenSourceMock {
	return &tokenSourceMock{shouldReturnError: shouldReturnError}
}

// TokenSourceMock is a local source of tokens, to be used for testing.
type tokenSourceMock struct {
	shouldReturnError bool
}

// Token returns a dummy token for testing.
// If shouldReturnError is true an error will be returned.
func (mock *tokenSourceMock) Token() (*oauth2.Token, error) {
	if mock.shouldReturnError {
		return nil, errors.New("An Error has occurred.")
	}
	return &oauth2.Token{
		AccessToken:  "<ACCESS_TOKEN>",
		TokenType:    "Bearer",
		RefreshToken: "<REFRESH_TOKEN>",
		Expiry:       time.Now().Add(30 * time.Minute),
	}, nil
}

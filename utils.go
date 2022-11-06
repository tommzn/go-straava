package strava

import (
	"golang.org/x/oauth2"
)

// TokenSourceFromAuthorizationCode is a helper to get a token source for an athorization code.
// HAve a look at to get more details how to obtain an authorization code for your app.
// Please persist returned refresh token, because an auth code can be used only once.
func TokenSourceFromAuthorizationCode(oauthConfig oauth2.Config, authCode string) (oauth2.TokenSource, error) {
	token, err := oauthConfig.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		return nil, err
	}
	return oauthConfig.TokenSource(oauth2.NoContext, token), nil
}

// TokenSourceFromRefreshToken is a helper to create a token source for an existing refresh token.
func TokenSourceFromRefreshToken(oauthConfig oauth2.Config, refreshToken string) (oauth2.TokenSource, error) {
	return oauthConfig.TokenSource(oauth2.NoContext, &oauth2.Token{RefreshToken: refreshToken}), nil
}

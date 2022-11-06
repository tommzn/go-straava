[![Go Reference](https://pkg.go.dev/badge/github.com/tommzn/go-config.svg)](https://pkg.go.dev/github.com/tommzn/go-strava)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tommzn/go-strava)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/tommzn/go-strava)
[![Go Report Card](https://goreportcard.com/badge/github.com/tommzn/go-strava)](https://goreportcard.com/report/github.com/tommzn/go-strava)
[![Actions Status](https://github.com/tommzn/go-strava/actions/workflows/go.pkg.auto-ci.yml/badge.svg)](https://github.com/tommzn/go-strava/actions)
![Strava API Version](https://img.shields.io/badge/Strava_API-v3-blue)

# Strava API Client
A client to access APIs provided by Strava.

## API Version
This client is implemented for Strave API v3.

# Authtication
Strave uses OAuth2 for their APIs. This client uses [TokenSource](https://pkg.go.dev/golang.org/x/oauth2#TokenSource) from [oauth2](https://pkg.go.dev/golang.org/x/oauth2) for authentication. 
Have a look at [Getting Started with the Strava API](https://developers.strava.com/docs/getting-started/) to get details about how to authenticate your app to access the API.
Base of OAuth2 authentication is [oauth2.Config](https://pkg.go.dev/golang.org/x/oauth2#Config) where you specifiy your client id, client secret and scopes.

## Helper
There're to helper to create required token source. Both methods expect a [oauth2.Config](https://pkg.go.dev/golang.org/x/oauth2#Config).
### TokenSourceFromAuthorizationCode
This method uses [Exchange](https://pkg.go.dev/golang.org/x/oauth2#Config.Exchange) method on oauth2 config to get access and refresh token for an authorization code. You should persist your refresh token somewhere, because an authorization code can be used only once.
### TokenSourceFromRefreshToken
If you already have a refresh token you can use this method to create a token source for it.

# What about all the other models and endpoints?
Scope of this client is limited to models and endpoints I'm using in my project atm. Feel free to create an enhancement issue to request an extension.

# Links
- [Getting Started with the Strava API](https://developers.strava.com/docs/getting-started/)
- [Strava API Documentation](https://developers.strava.com/docs/reference/)

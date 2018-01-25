JWTKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/jwtkit)](https://goreportcard.com/report/github.com/gokit/jwtkit)

JWTKit implements a code generator which automatically generates go packages for jwt implementations for annotated structs.
It focuses around generating API code for a giving struct which would contain all data to be attached to a giving [jwt.Claims](https://godoc.org/github.com/dgrijalva/jwt-go#Claims).

It generates all necessary code to issue jwt access and refresh tokens for use in your application.

## Install

```
go get -u github.com/gokit/jwtkit
```

## Examples

See [Examples](./examples) for demonstrations of jwtkit code generation.

## Usage

Running the following commands instantly generates all necessary files and packages for giving code gen.

```go
> jwit generate
```

## How It works

### Struct Annotation

You annotate any giving struct with `@mongoapi` which marks giving struct has a target for code generation. 

Sample below:

```go
// UserClaim embodies data to be attached with a jwt claim.
// @jwt(Contract => UserContract)
type UserClaim struct {
	User User `json:"user"`
	Roles map[string]string `json:"roles"`
	Permissions map[string]string `json:"permissions"`
}

// UserContract embodies the data needed for authenticating a user.
type UserContract struct{
	Username string `json:"user_name"`
	Password string `json:"password"`
}
```

JWTKit generates specific interfaces to allow extensibility and also 
what underline backend technology would be used for either storing jwt
records and blacklist/active refresh tokens and associated records.

The JWTKit expects that the provided annotation must have a `Contract`
attribute pointing to a local exported struct which contains expected 
data to be received to authenticate a login session by the user to 
authenticate itself for access to the jwt authorization system.

```go

blacklist := mock.TokenBackend()
whitelist := mock.TokenBackend()
idb := mock.IdentityBackend()

jwter := userclaimjwt.NewJWTIdentity(userclaimjwt.JWTConfig{
	Maker:               noSecureUser,
	Signer:              "wellington",
	SigningSecret:       "All we want is to sign this",
	Method:              jwt.SigningMethodHS256,
	AccessTokenExpires:  700 * time.Millisecond,
	RefreshTokenExpires: 4 * time.Second,
}, blacklist, whitelist, idb)

var cred pkg.CreateUserSession
if jsnerr := json.Unmarshal([]byte(contractDataJSON), &cred); jsnerr != nil {
	panic(jsnerr)
}

auth, err := jwter.Grant(context.Background(), cred)
if err != nil {
	panic(jsnerr)
}

```

Where `auth` is:
```json
{
	"expires": 323231,
	"refresh_token":"",
	"access_token":"",
	"token_type": "Bearer",
	"refresh_expires": 323231,
}
```

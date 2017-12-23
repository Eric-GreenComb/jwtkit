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

## Usage

Running the following commands instantly generates all necessary files and packages for giving code gen.

```go
> jwit generate
```

## How It works

### Struct Annotation

You annotate any giving struct with `@mongoapi` which marks giving struct has a target for code generation. 

*All struct must have a `PublicID` field.*

Sample below:

```go
// UserClaim embodies data to be attached with a jwt claim.
// @jwt
type UserClaim struct {
	User User `json:"user"`
	Roles map[string]string `json:"roles"`
	Permissions map[string]string `json:"permissions"`
}
```

JWTkit generates specific interfaces to allow extensibility and also 
what underline backend technology would be used for either storing jwt
records and blacklist/active refresh tokens and associated records.

Each interface is included with all generated files.

Sample interfaces to be implemented for `User` struct (only one set is needed):

```go
```


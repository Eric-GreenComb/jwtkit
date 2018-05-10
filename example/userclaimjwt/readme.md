UserClaim HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com\gokit\jwtkit\example\userclaimjwt)](https://goreportcard.com/report/github.com\gokit\jwtkit\example\userclaimjwt)

UserClaim HTTP API is a auto generated http api for the struct `UserClaim`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type IdentityBackend interface {
	Count(context.Context) (int, error)
	Delete(context.Context, string) error
	Get(context.Context, string) (Identity, error)
	Update(context.Context, string, Identity) error
	Revoke(context.Context, string) error
	Refresh(context.Context, string) (JWTAuth, error)
	GetAll(context.Context, string, string, int, int) ([]Identity, int, error)
	Grant(context.Context, example.CreateUserSession) (JWTAuth, error)
	Authenticate(context.Context, string) (example.UserClaim, error)
}
```


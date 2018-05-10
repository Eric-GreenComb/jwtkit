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


```go

var contractDataJSON = `{


    "password":	"3mkbzeb1pooel6mxw38z",

    "expiration":	null,

    "email":	"DorothyBurns@Buzzbean.gov"

}`

blacklist := mock.TokenBackend()
whitelist := mock.TokenBackend()
idb := mock.IdentityBackend()
jwter := userclaimjwt.NewJWTIdentity(userclaimjwt.JWTConfig{
    Maker:               noSecureUser,
    Signer:              "wellington",
    Secrets:             secretFunc,
    Method:              jwt.SigningMethodHS256,
    AccessTokenExpires:  600 * time.Millisecond,
    RefreshTokenExpires: 1 * time.Second,
}, blacklist, whitelist, idb)

var cred example.CreateUserSession
if jsnerr := json.Unmarshal([]byte(contractDataJSON), &cred); jsnerr != nil {
    log.Fatal(jserr)
}

auth, err := jwter.Grant(context.Background(), cred)
```
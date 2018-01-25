package example

import "time"

//@jwt(Contract => CreateUserSession)
type UserClaim struct {
	User  User     `json:"user"`
	Roles []string `json:"roles"`
}

type User struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	PublicID      string `json:"public_id"`
	TenantID      string `json:"tenant_id"`
	PrivateID     string `json:"private_id"`
	TwoFactorAuth bool   `json:"two_factor_auth"`
}

type CreateUserSession struct {
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	Expiration time.Duration `json:"expiration"`
}

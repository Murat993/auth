package dto

import (
	desc "github.com/Murat993/auth/pkg/user_v1"
	"github.com/dgrijalva/jwt-go"
)

const (
	ExamplePath = "/note_v1.NoteV1/Get"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string    `json:"username"`
	Role     desc.Role `json:"role"`
}

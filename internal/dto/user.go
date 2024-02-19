package dto

import (
	"database/sql"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/wrappers"
	"time"
)

type UserCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            desc.Role
}

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      desc.Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserUpdate struct {
	ID    string
	Name  *wrappers.StringValue
	Email *wrappers.StringValue
}

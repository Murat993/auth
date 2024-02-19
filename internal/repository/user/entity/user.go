package entity

import (
	"database/sql"
	desc "github.com/Murat993/auth/pkg/user_v1"
	"time"
)

type User struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      desc.Role    `db:"role"`
	CreatedAt time.Time    `db:"createdAt"`
	UpdatedAt sql.NullTime `db:"updatedAt"`
}

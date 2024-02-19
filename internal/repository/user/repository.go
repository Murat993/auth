package user

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/Murat993/auth/internal/client/db"
	"github.com/Murat993/auth/internal/dto"
	def "github.com/Murat993/auth/internal/repository"
	"github.com/Murat993/auth/internal/repository/user/converter"
	modelRepo "github.com/Murat993/auth/internal/repository/user/entity"
)

var _ def.UserRepository = (*repo)(nil)

const (
	tableName = "note"

	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	roleColumn            = "role"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) *repo {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, userCreate *dto.UserCreate) (string, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, passwordConfirmColumn, roleColumn).
		Values(userCreate.Name, userCreate.Email, userCreate.Password, userCreate.PasswordConfirm, userCreate.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id string) (*dto.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, userUpdate *dto.UserUpdate) (string, error) {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, userUpdate.Name).
		Set(emailColumn, userUpdate.Email).
		Where(sq.Eq{"id": userUpdate.ID}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return fmt.Errorf("failed to Delete user: %v, tag: %v", err, res)
	}

	return nil
}

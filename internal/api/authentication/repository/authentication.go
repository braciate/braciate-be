package authRepository

import (
	"database/sql"
	"errors"
	"github.com/braciate/braciate-be/internal/api/authentication"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type UserDB struct {
	ID           string          `db:"id"`
	Username     string          `db:"name"`
	Password     string          `db:"password"`
	NIM          string          `db:"nim"`
	Email        string          `db:"email"`
	Faculty      string          `db:"faculty"`
	StudyProgram string          `db:"study_program"`
	Role         entity.UserRole `db:"role"`
}

func (user *UserDB) format() entity.User {
	return entity.User{
		ID:           user.ID,
		Username:     user.Username,
		NIM:          user.NIM,
		Email:        user.Email,
		Faculty:      user.Faculty,
		StudyProgram: user.StudyProgram,
		Role:         user.Role,
		Password:     user.Password,
	}
}

func (r *AuthRepository) GetUserByEmailOrNIM(ctx context.Context, identifier string) (entity.User, error) {
	argsKV := map[string]interface{}{
		"identifier": identifier,
	}

	query, args, err := sqlx.Named(queryGetUserByEmailOrNIM, argsKV)
	if err != nil {
		return entity.User{}, err
	}
	query = r.q.Rebind(query)

	var user UserDB
	if err := r.q.QueryRowxContext(ctx, query, args...).StructScan(&user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, authentication.ErrRecordNotFound
		}
		return entity.User{}, err
	}

	return user.format(), nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user entity.User) (string, error) {
	argsKV := map[string]interface{}{
		"id":            user.ID,
		"name":          user.Username,
		"nim":           user.NIM,
		"email":         user.Email,
		"faculty":       user.Faculty,
		"study_program": user.StudyProgram,
		"role":          user.Role,
		"password":      user.Password,
	}

	query, args, err := sqlx.Named(queryCreateUser, argsKV)
	if err != nil {
		return "", err
	}
	query = r.q.Rebind(query)

	var id string
	if err := r.q.QueryRowxContext(ctx, query, args...).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

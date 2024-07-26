package user

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	repository "github.com/lookandhate/course_auth/internal/repository/model"
	"github.com/lookandhate/course_auth/internal/service/model"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

const (
	userTable = "users"

	idColumn = "id"

	emailColumn        = "email"
	passwordHashColumn = "password_hash"
	nameColumn         = "name"
	roleColumn         = "role"

	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

// NewPostgresRepository creates PostgresRepository instance.
func NewPostgresRepository(connectionPool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pgx: connectionPool}
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user *repository.CreateUserModel) (int, error) {
	builder := squirrel.Insert(userTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(emailColumn, passwordHashColumn, nameColumn, roleColumn).
		Values(user.Email, user.Password, user.Name, user.Role).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *PostgresRepository) GetUser(ctx context.Context, id int) (*model.UserModel, error) {
	builder := squirrel.
		Select(idColumn, emailColumn, passwordHashColumn, nameColumn, roleColumn).
		PlaceholderFormat(squirrel.Dollar).
		From(userTable).
		Where(squirrel.Eq{idColumn: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.UserModel

	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, user *model.UpdateUserModel) (*model.UserModel, error) {
	builder := squirrel.Update(userTable).PlaceholderFormat(squirrel.Dollar).Where(squirrel.Eq{idColumn: user.ID})

	if user.Password != nil {
		builder = builder.Set(passwordHashColumn, user.Password)
	}
	if user.Name != nil {
		builder = builder.Set(nameColumn, user.Name)
	}
	if user.Role != int(model.UserUnknownRole) {
		builder = builder.Set(roleColumn, user.Role)
	}
	if user.Email != nil {
		builder = builder.Set(emailColumn, user.Email)
	}

	builder = builder.
		Set(updatedAtColumn, time.Now()).
		Suffix(
			fmt.Sprintf(
				"RETURNING %s, %s, %s, %s, %s, %s",
				idColumn, emailColumn, nameColumn, roleColumn, createdAtColumn, updatedAtColumn),
		)

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	updatedUser := &model.UserModel{}
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(
		&updatedUser.ID,
		&updatedUser.Email,
		&updatedUser.Name,
		&updatedUser.Role,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id int) error {
	builder := squirrel.Delete(userTable).Where(squirrel.Eq{idColumn: id})
	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	log.Println(sql, args)
	_, err = r.pgx.Exec(ctx, sql, args...)
	return err
}

func (r *PostgresRepository) CheckUserExists(ctx context.Context, id int) (bool, error) {
	var exists bool

	builder := squirrel.Select(
		fmt.Sprintf("EXISTS(SELECT 1 FROM %s WHERE id = %s) AS user_exists", userTable, strconv.Itoa(id)),
	)
	sql, args, err := builder.ToSql()

	if err != nil {
		return false, err
	}
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&exists)
	return exists, err
}

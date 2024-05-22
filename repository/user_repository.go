package repository

import (
	"context"
	"errors"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/entity"
	"github.com/caturarp/laporplat/query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	FindUserByEmail(context.Context, string) (entity.User, error)
	FindUserByID(context.Context, uint) (*dto.UserDetailResponse, error)
	AddNewUser(context.Context, *dto.RegisterRequest, pgx.Tx) (*dto.IDResponse, error)
	UpdateUser(context.Context, *dto.UpdateUserRequest) (*dto.IDResponse, error)
	UpdateName(context.Context, string, uint, pgx.Tx) (*dto.IDResponse, error)
	ListUser(context.Context) ([]dto.UserDetailResponse, error)
}
type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) ListUser(ctx context.Context) ([]dto.UserDetailResponse, error) {
	var users []dto.UserDetailResponse
	q := `SELECT name, email FROM users WHERE deleted_at IS NULL`
	rows, err := u.db.Query(ctx, q)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []dto.UserDetailResponse{}, err
		}
		return users, apperr.ErrUserNotFound
	}
	defer rows.Close()

	for rows.Next() {
		var user dto.UserDetailResponse
		if err := rows.Scan(&user.Name, &user.Email); err != nil {
			return []dto.UserDetailResponse{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := u.db.QueryRow(ctx, query.FindUserByEmail, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, err
		}
		return user, apperr.ErrFindUserByEmail
	}
	return user, nil
}

func (u *userRepository) FindUserByID(ctx context.Context, userID uint) (*dto.UserDetailResponse, error) {
	user := dto.UserDetailResponse{}
	err := u.db.QueryRow(ctx, query.FindUserByID, userID).Scan(&user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrUserNotFound
		}
		return nil, apperr.ErrFindUserByIDQuery
	}
	return &user, nil
}

func (u *userRepository) AddNewUser(ctx context.Context, userInfo *dto.RegisterRequest, tx pgx.Tx) (*dto.IDResponse, error) {
	resp := dto.IDResponse{}
	err := tx.QueryRow(ctx, query.AddNewUser, userInfo.Email, userInfo.Password, userInfo.VerifiedAt).Scan(&resp.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrRecordNotFound
		}
		return nil, apperr.ErrNewUserQuery
	}

	return &resp, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, userDetail *dto.UpdateUserRequest) (*dto.IDResponse, error) {
	resp := dto.IDResponse{}
	err := u.db.QueryRow(ctx, query.UpdateUser, userDetail.Name, userDetail.ID).Scan(&resp.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrRecordNotFound
		}
		return nil, apperr.ErrUpdateUser
	}

	return &resp, nil
}

func (u *userRepository) UpdateName(ctx context.Context, name string, id uint, tx pgx.Tx) (*dto.IDResponse, error) {
	var userID uint
	err := tx.QueryRow(ctx, query.UpdateUserName, name, id).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrRecordNotFound
		}
		return nil, apperr.ErrUpdateUser
	}

	return &dto.IDResponse{ID: userID}, nil
}

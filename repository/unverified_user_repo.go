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

type UnverifiedUserRepository interface {
	FindUnverifiedUser(context.Context, dto.VerifyRequest) error
	AddUnverifiedUser(context.Context, entity.UnverifiedUser) (*dto.VerifyResponse, error)
	DeleteUnverifiedUser(context.Context, dto.DeleteUnverifiedUserRequest, pgx.Tx) (*dto.IDResponse, error)
}
type unverifiedUserRepository struct {
	db *pgxpool.Pool
}

func NewUnverifiedUserRepository(db *pgxpool.Pool) UnverifiedUserRepository {
	return &unverifiedUserRepository{
		db: db,
	}
}

func (v *unverifiedUserRepository) FindUnverifiedUser(ctx context.Context, userInfo dto.VerifyRequest) error {
	_, err := v.db.Exec(ctx, query.FindUnverifiedUser, userInfo.Email, userInfo.Code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return apperr.ErrEmailUnregistered
		}
		return err
	}
	return nil
}

func (v *unverifiedUserRepository) AddUnverifiedUser(ctx context.Context, unverifiedUser entity.UnverifiedUser) (*dto.VerifyResponse, error) {
	var response dto.VerifyResponse

	tx, err := v.db.Begin(ctx)
	if err != nil {
		return &response, err
	}
	defer tx.Rollback(ctx)

	result, err := tx.Exec(ctx, query.FindUnverifiedUserByEmail, unverifiedUser.Email)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return &response, err
		}
	}

	if result.RowsAffected() > 0 {
		_, err := tx.Exec(ctx, query.DeleteUnverifiedUser, unverifiedUser.Email)
		if err != nil {
			return &response, apperr.ErrDeleteVerifyRequest
		}
	}

	var code string
	err = tx.QueryRow(ctx, query.AddUnverifiedUser, unverifiedUser.Email, unverifiedUser.Code, unverifiedUser.ExpiredAt).Scan(&code)
	if err != nil {
		return &response, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return &response, apperr.ErrTxCommit
	}
	response.Code = code
	return &response, nil
}

func (v *unverifiedUserRepository) DeleteUnverifiedUser(ctx context.Context, userInfo dto.DeleteUnverifiedUserRequest, tx pgx.Tx) (*dto.IDResponse, error) {
	resp := dto.IDResponse{}
	err := tx.QueryRow(ctx, query.DeleteUnverifiedUser, userInfo.Email).Scan(&resp.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrRecordNotFound
		}
		return nil, apperr.ErrDeleteVerifyRequest
	}

	return &resp, nil
}

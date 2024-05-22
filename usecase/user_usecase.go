package usecase

import (
	"context"
	"errors"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/entity"
	"github.com/caturarp/laporplat/repository"
	"github.com/jackc/pgx/v5"
)

type UserUsecase interface {
	ListUser(context.Context) ([]dto.UserDetailResponse, error)
	FindUser(context.Context, dto.UserParameter) (*entity.User, error)
	GetUserDetail(ctx context.Context) (*dto.UserDetailResponse, error)
	UpdateUser(ctx context.Context, userDetail *dto.UpdateUserRequest) (*dto.IDResponse, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: r,
	}
}

func (u *userUsecase) ListUser(ctx context.Context) ([]dto.UserDetailResponse, error) {
	users, err := u.userRepository.ListUser(ctx)
	if err != nil {
		return []dto.UserDetailResponse{}, err
	}
	return users, nil
}

func (u *userUsecase) FindUser(ctx context.Context, param dto.UserParameter) (*entity.User, error) {
	user, err := u.userRepository.FindUserByEmail(ctx, param.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrUserNotFound
		} else {
			return nil, apperr.ErrFindUserByEmail
		}
	}
	return &user, nil
}

func (u *userUsecase) GetUserDetail(ctx context.Context) (*dto.UserDetailResponse, error) {
	userID := ctx.Value("user_id").(uint)
	return u.userRepository.FindUserByID(ctx, userID)
}

func (u *userUsecase) UpdateUser(ctx context.Context, userDetail *dto.UpdateUserRequest) (*dto.IDResponse, error) {
	return u.userRepository.UpdateUser(ctx, userDetail)
}

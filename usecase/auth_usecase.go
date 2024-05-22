package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/entity"
	"github.com/caturarp/laporplat/external/mail"
	"github.com/caturarp/laporplat/repository"
	"github.com/caturarp/laporplat/util"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)
	RequestVerification(context.Context, dto.VerifyRequest, uint) (*dto.VerifyResponse, error)
	CompleteUserRegistration(context.Context, dto.RegisterRequest) (*dto.IDResponse, error)
}

type authUsecase struct {
	userRepository           repository.UserRepository
	unverifiedUserRepository repository.UnverifiedUserRepository
	smtpRepository           mail.SMTPMailer
	db                       *pgxpool.Pool
}

func NewAuthUsecase(ur repository.UserRepository, uur repository.UnverifiedUserRepository, sr mail.SMTPMailer, db *pgxpool.Pool) AuthUsecase {
	return &authUsecase{
		userRepository:           ur,
		unverifiedUserRepository: uur,
		smtpRepository:           sr,
		db:                       db,
	}
}

func (a *authUsecase) Login(ctx context.Context, loginInfo dto.LoginRequest) (*dto.LoginResponse, error) {
	var response dto.LoginResponse
	var token string

	user, err := a.userRepository.FindUserByEmail(ctx, loginInfo.Email)
	if err != nil {
		return nil, err
	}
	if !util.ComparePassword(user.Password, loginInfo.Password) {
		return nil, apperr.ErrWrongCredentials
	}
	var personalID uint

	token, _ = dto.GenerateAccessToken(dto.JWTClaims{
		UserID:     user.ID,
		PersonalID: personalID,
	})
	response.AccessToken = token
	return &response, nil
}

func (a *authUsecase) RequestVerification(ctx context.Context, registerInfo dto.VerifyRequest, role uint) (*dto.VerifyResponse, error) {
	user, err := a.userRepository.FindUserByEmail(ctx, registerInfo.Email)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
	}
	if !user.IsEmpty() {
		return nil, apperr.ErrEmailALreadyUsed
	}

	code := util.GenerateCode(registerInfo)
	unverifiedUser := entity.UnverifiedUser{
		Email:     registerInfo.Email,
		Code:      code,
		ExpiredAt: time.Now().Add(5 * time.Minute),
	}
	verifyResponse, err := a.unverifiedUserRepository.AddUnverifiedUser(ctx, unverifiedUser)
	if err != nil {
		return nil, err
	}

	sendEmailInfo := dto.SendVerificationMailRequest{}

	verificationLink := fmt.Sprintf("http://%s/auth/verify/?email=%s&code=%s", os.Getenv("APP_HOST"), registerInfo.Email, code)
	sendEmailInfo = dto.SendVerificationMailRequest{
		EmailRecipients: []string{unverifiedUser.Email},
		Subject:         "Account Activation Request | APP_NAME",
		Content:         verificationLink,
	}

	err = a.smtpRepository.SendMail(ctx, sendEmailInfo)
	if err != nil {
		return nil, err
	}

	return verifyResponse, nil
}

func (a *authUsecase) CompleteUserRegistration(ctx context.Context, completeRegisterInfo dto.RegisterRequest) (*dto.IDResponse, error) {
	completeRegisterInfo.VerifiedAt = time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(completeRegisterInfo.Password), 10)
	if err != nil {
		return nil, apperr.ErrGenerateHashPassword
	}
	completeRegisterInfo.Password = string(hashedPassword)

	err = a.unverifiedUserRepository.FindUnverifiedUser(ctx, dto.VerifyRequest{Email: completeRegisterInfo.Email, Code: completeRegisterInfo.Code})
	if err != nil {
		return nil, err
	}

	tx, err := a.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	_, err = a.unverifiedUserRepository.DeleteUnverifiedUser(ctx, dto.TransformInfoToDeleteRequest(completeRegisterInfo.Email), tx)
	if err != nil {
		return nil, err
	}

	createdUserID, err := a.userRepository.AddNewUser(ctx, &completeRegisterInfo, tx)
	if err != nil {
		return nil, apperr.ErrNewUserQuery
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, apperr.ErrTxCommit
	}

	return createdUserID, nil
}

package handler

import (
	"errors"
	"net/http"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/common"
	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(au usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: au,
	}
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var loginInfo dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginInfo)
	if err != nil {
		ctx.Error(apperr.ErrInvalidBody)
		return
	}

	tokenResponse, err := a.authUsecase.Login(ctx, loginInfo)
	if err != nil {
		if errors.Is(err, apperr.ErrUserNotFound) {
			ctx.Error(apperr.ErrWrongCredentials)
			return
		}
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Data: tokenResponse})
}

func (a *AuthHandler) RequestRegister(ctx *gin.Context) {
	var requestInfo dto.VerifyRequest
	err := ctx.ShouldBindJSON(&requestInfo)
	if err != nil {
		ctx.Error(apperr.ErrInvalidBody)
		return
	}
	unverifiedUser, err := a.authUsecase.RequestVerification(ctx, requestInfo, common.RoleUser)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{Data: unverifiedUser})
}

func (a *AuthHandler) VerifyRegister(ctx *gin.Context) {
	var verificationInfo dto.RegisterRequest
	email := ctx.Query("email")
	code := ctx.Query("code")

	verificationInfo.Email = email
	verificationInfo.Code = code

	err := ctx.ShouldBindJSON(&verificationInfo)
	if err != nil {
		ctx.Error(apperr.ErrInvalidBody)
		return
	}
	userCreated, err := a.authUsecase.CompleteUserRegistration(ctx, verificationInfo)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{Data: userCreated})
}

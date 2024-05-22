package handler

import (
	"net/http"

	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUsecase
}

func NewUserHandler(userUseCase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (u *UserHandler) ListUser(ctx *gin.Context) {
	users, err := u.userUseCase.ListUser(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{Data: users})
}

func (u *UserHandler) FindUser(ctx *gin.Context) {
	var req dto.UserParameter
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	user, err := u.userUseCase.FindUser(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: user})
}

func (u *UserHandler) GetUserDetail(ctx *gin.Context) {
	user, err := u.userUseCase.GetUserDetail(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: user})
}

func (u *UserHandler) UpdateUserDetail(ctx *gin.Context) {
	req := dto.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	userID := ctx.Value("user_id").(uint)
	req.ID = userID

	userDetailResponse, err := u.userUseCase.UpdateUser(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Data: userDetailResponse})
}

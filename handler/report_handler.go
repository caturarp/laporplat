package handler

import (
	"net/http"
	"strconv"

	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/usecase"
	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportUseCase usecase.ReportUsecase
}

func NewReportHandler(reportUseCase usecase.ReportUsecase) *ReportHandler {
	return &ReportHandler{
		reportUseCase: reportUseCase,
	}
}

func (r *ReportHandler) ListReport(ctx *gin.Context) {
	reports, err := r.reportUseCase.ListReport(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, reports)
}

func (r *ReportHandler) FindReportByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}
	report, err := r.reportUseCase.FindReportByID(ctx, uint(idUint))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, report)
}

func (r *ReportHandler) CreateReport(ctx *gin.Context) {
	var req dto.CreateReportRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = r.reportUseCase.CreateReport(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (r *ReportHandler) UpdateReport(ctx *gin.Context) {
	var req dto.UpdateReportRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = r.reportUseCase.UpdateReport(ctx, &req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (r *ReportHandler) DeleteReport(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = r.reportUseCase.DeleteReport(ctx, uint(idUint))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

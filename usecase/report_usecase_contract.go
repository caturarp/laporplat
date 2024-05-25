package usecase

import (
	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/entity"
	"github.com/gin-gonic/gin"
)

type ReportUsecase interface {
	ListReport(ctx *gin.Context) ([]entity.Report, error)
	FindReportByID(ctx *gin.Context, id uint) (*entity.Report, error)
	CreateReport(ctx *gin.Context, req *dto.CreateReportRequest) error
	UpdateReport(ctx *gin.Context, req *dto.UpdateReportRequest) error
	DeleteReport(ctx *gin.Context, id uint) error
}

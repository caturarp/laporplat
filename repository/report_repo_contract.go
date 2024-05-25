package repository

import (
	"context"

	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/entity"
	"github.com/jackc/pgx/v5"
)

type ReportRepository interface {
	FindReportByID(context.Context, uint) (*entity.Report, error)
	CreateReport(context.Context, pgx.Tx, *dto.CreateReportRequest) error
	SaveReport(context.Context, pgx.Tx, *dto.UpdateReportRequest) error
	DeleteReport(context.Context, pgx.Tx, uint) error
	ListReportByLicensePlate(context.Context, string) ([]entity.Report, error)
	ListReportByDriverName(context.Context, string) ([]entity.Report, error)
}

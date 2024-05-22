package repository

import (
	"context"

	"github.com/caturarp/laporplat/entity"
)

type ReportRepository interface {
	FindReportByID(context.Context, entity.Report, uint) error
}

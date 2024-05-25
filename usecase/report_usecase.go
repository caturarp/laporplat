package usecase

import (
	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/entity"
	"github.com/caturarp/laporplat/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type reportUsecase struct {
	reportRepo repository.ReportRepository
	pool       *pgxpool.Pool
}

func NewReportUsecase(reportRepo repository.ReportRepository, pool *pgxpool.Pool) ReportUsecase {
	return &reportUsecase{
		reportRepo: reportRepo,
		pool:       pool,
	}
}

func (r *reportUsecase) ListReport(ctx *gin.Context) ([]entity.Report, error) {
	return r.reportRepo.ListReportByDriverName(ctx, "string")
}
func (r *reportUsecase) FindReportByID(ctx *gin.Context, id uint) (*entity.Report, error) {
	return r.reportRepo.FindReportByID(ctx, id)
}
func (r *reportUsecase) CreateReport(ctx *gin.Context, req *dto.CreateReportRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return apperr.ErrDatabaseConnection
	}
	defer tx.Rollback(ctx)
	err = r.reportRepo.CreateReport(ctx, tx, req)
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return apperr.ErrTxCommit
	}
	return nil
}
func (r *reportUsecase) UpdateReport(ctx *gin.Context, req *dto.UpdateReportRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return apperr.ErrDatabaseConnection
	}
	defer tx.Rollback(ctx)
	err = r.reportRepo.SaveReport(ctx, tx, req)
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return apperr.ErrTxCommit
	}
	return nil
}
func (r *reportUsecase) DeleteReport(ctx *gin.Context, id uint) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return apperr.ErrDatabaseConnection
	}
	defer tx.Rollback(ctx)
	err = r.reportRepo.DeleteReport(ctx, tx, id)
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return apperr.ErrTxCommit
	}
	return nil
}

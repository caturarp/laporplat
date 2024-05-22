package repository

import (
	"context"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/entity"
	"github.com/jackc/pgx/v5"
)

func findReportByID(ctx context.Context, tx pgx.Tx, id uint) (*entity.Report, error) {
	q := `SELECT 
  id,
  vehicle_id,
  license_plate,
  report_type,
  description,
  created_at
  WHERE
  id = $1`

	report := new(entity.Report)
	err := tx.QueryRow(ctx, q, id).Scan(&report.ID, &report.VehicleID, &report.LicensePlate, &report.ReportType, &report.Description, &report.CreatedAt)
	if err != nil {
		return nil, apperr.ErrDatabaseQuery
	}
	return report, nil
}

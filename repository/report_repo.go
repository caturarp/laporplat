package repository

import (
	"context"
	"errors"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/dto"
	"github.com/caturarp/laporplat/entity"
	"github.com/caturarp/laporplat/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type reportRepository struct {
	db *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool) ReportRepository {
	return &reportRepository{
		db: db,
	}
}
func (r *reportRepository) FindReportByID(ctx context.Context, id uint) (*entity.Report, error) {
	q := `SELECT 
		id,
		vehicle_id,
		license_plate,
		report_type,
		description,
		created_at
		WHERE
		id = $1 AND
		deleted_at IS NULL`

	report := new(entity.Report)
	err := r.db.QueryRow(ctx, q, id).Scan(&report.ID, &report.VehicleID, &report.LicensePlate, &report.ReportType, &report.Description, &report.CreatedAt)
	if err != nil {
		return nil, apperr.ErrDatabaseQuery
	}
	return report, nil
}
func (r *reportRepository) CreateReport(ctx context.Context, tx pgx.Tx, report *dto.CreateReportRequest) error {
	q := `INSERT INTO reports (vehicle_id, license_plate, report_type, description) VALUES ($1, $2, $3, $4) RETURNING id`
	row, err := tx.Exec(ctx, q, report.VehicleID, report.LicensePlate, report.Type, report.Description)
	if err != nil {
		return apperr.ErrDatabaseQuery
	}
	if row.RowsAffected() == 0 {
		return apperr.ErrDatabaseQuery
	}
	return nil
}
func (r *reportRepository) SaveReport(ctx context.Context, tx pgx.Tx, req *dto.UpdateReportRequest) error {
	q := `UPDATE reports SET status = 'saved' WHERE id = $1`
	_, err := tx.Exec(ctx, q, req.ID)
	if err != nil {
		return apperr.ErrDatabaseQuery
	}
	return nil
}
func (r *reportRepository) DeleteReport(ctx context.Context, tx pgx.Tx, id uint) error {
	q := `UPDATE reports SET deleted_at = NOW() WHERE id = $1`
	_, err := tx.Exec(ctx, q, id)
	if err != nil {
		return apperr.ErrDatabaseQuery
	}
	return nil
}
func (r *reportRepository) ListReportByLicensePlate(ctx context.Context, licensePlate string) ([]entity.Report, error) {
	q := `SELECT * FROM reports where license_plate = $1 AND deleted_at IS NULL`
	rows, err := r.db.Query(ctx, q, licensePlate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperr.ErrRecordNotFound
		}
		return nil, apperr.ErrDatabaseQuery
	}
	defer rows.Close()
	var reports []entity.Report
	for rows.Next() {
		var report entity.Report
		if err := util.ScanStruct(rows, &report); err != nil { // Notice the '&' to pass a pointer
			return nil, apperr.ErrDatabaseQuery
		}
		reports = append(reports, report)

	}
	return reports, nil
}

func (r *reportRepository) ListReportByDriverName(ctx context.Context, name string) ([]entity.Report, error) {
	q := `SELECT * FROM reports where driver_name = $1 AND deleted_at IS NULL`
	rows, err := r.db.Query(ctx, q, name)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, apperr.ErrRecordNotFound
	}
	if err != nil {
		return nil, apperr.ErrDatabaseQuery
	}

	defer rows.Close()
	var reports []entity.Report
	for rows.Next() {
		var report entity.Report
		if err := util.ScanStruct(rows, &report); err != nil {
			return nil, apperr.ErrDatabaseQuery
		}
		reports = append(reports, report)
	}
	return reports, nil
}

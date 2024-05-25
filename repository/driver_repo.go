package repository

import (
	"context"
	"errors"

	"github.com/caturarp/laporplat/apperr"
	"github.com/caturarp/laporplat/entity"
	"github.com/caturarp/laporplat/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type driverRepository struct {
	db *pgxpool.Pool
}

func NewDriverRepository(db *pgxpool.Pool) DriverRepository {
	return &driverRepository{db: db}
}

func (r *driverRepository) FindDriverByID(ctx context.Context, id uint) (*entity.Driver, error) {
	driver := new(entity.Driver)
	err := r.db.QueryRow(ctx, "SELECT * FROM drivers WHERE id = $1", id).Scan(
		&driver.ID, &driver.Name, &driver.LicenseNumber, &driver.RideService, &driver.PhoneNumber, &driver.Email,
	)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, apperr.ErrRecordNotFound
	}
	if err != nil {
		return nil, apperr.ErrDatabaseQuery
	}

	return driver, nil
}

func (r *driverRepository) CreateDriver(ctx context.Context, driver entity.Driver) error {
	q := "INSERT INTO drivers (name, license_number, ride_service, phone_number, email) VALUES ($1, $2, $3, $4, $5)"
	if _, err := r.db.Exec(ctx, q,
		driver.Name, driver.LicenseNumber, driver.RideService, driver.PhoneNumber, driver.Email,
	); err != nil {
		return apperr.ErrDatabaseQuery
	}
	return nil
}

func (r *driverRepository) SaveDriver(ctx context.Context, driver entity.Driver) error {
	if _, err := r.db.Exec(ctx,
		"UPDATE drivers SET name = $1, license_number = $2, ride_service = $3, phone_number = $4, email = $5 WHERE id = $6",
		driver.Name, driver.LicenseNumber, driver.RideService, driver.PhoneNumber, driver.Email, driver.ID,
	); err != nil {
		return apperr.ErrDatabaseQuery
	}

	return nil
}

func (r *driverRepository) DeleteDriver(ctx context.Context, id uint) error {
	if _, err := r.db.Exec(ctx, "DELETE FROM drivers WHERE id = $1", id); err != nil {
		return apperr.ErrDatabaseQuery
	}
	return nil
}

func (r *driverRepository) ListDriverByLicensePlate(ctx context.Context, licensePlate string) ([]entity.Driver, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM drivers WHERE license_number = $1", licensePlate)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, apperr.ErrRecordNotFound
	}
	if err != nil {

		return nil, apperr.ErrDatabaseQuery
	}
	var drivers []entity.Driver
	for rows.Next() {
		var driver entity.Driver
		if err := util.ScanStruct(rows, &driver); err != nil { // Notice the '&' to pass a pointer
			return nil, apperr.ErrDatabaseQuery
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

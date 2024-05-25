package repository

import (
	"context"

	"github.com/caturarp/laporplat/entity"
)

type DriverRepository interface {
	FindDriverByID(context.Context, uint) (*entity.Driver, error)
	CreateDriver(context.Context, entity.Driver) error
	SaveDriver(context.Context, entity.Driver) error
	DeleteDriver(context.Context, uint) error
	ListDriverByLicensePlate(context.Context, string) ([]entity.Driver, error)
}

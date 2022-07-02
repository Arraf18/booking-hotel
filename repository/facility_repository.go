package repository

import (
	"booking_hotel/model/domain"
	"context"
	"database/sql"
)

type FacilityRepository interface {
	Save(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility
	Update(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility
	Delete(ctx context.Context, tx *sql.Tx, facility domain.Facility)
	FindById(ctx context.Context, tx *sql.Tx, facilityId int) (domain.Facility, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Facility
}

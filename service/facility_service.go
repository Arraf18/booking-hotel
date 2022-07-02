package service

import (
	"booking_hotel/model/web"
	"context"
)

type FacilityService interface {
	Create(ctx context.Context, request web.FacilityCreateRequest) web.FacilityResponse
	Update(ctx context.Context, request web.FacilityUpdateRequest) web.FacilityResponse
	Delete(ctx context.Context, facilityId int)
	FindById(ctx context.Context, facilityId int) web.FacilityResponse
	FindAll(ctx context.Context) []web.FacilityResponse
}

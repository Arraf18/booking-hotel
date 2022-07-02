package service

import (
	"booking_hotel/model/web"
	"context"
)

type MeetingRoomService interface {
	Create(ctx context.Context, request web.MeetingRoomCreateRequest) web.MeetingRoomResponse
	Update(ctx context.Context, request web.MeetingRoomUpdateRequest) web.MeetingRoomResponse
	Delete(ctx context.Context, meeting_roomId int)
	FindById(ctx context.Context, meeting_roomId int) web.MeetingRoomResponse
	FindAll(ctx context.Context) []web.MeetingRoomResponse
}

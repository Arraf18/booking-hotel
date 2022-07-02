package service

import (
	"booking_hotel/exception"
	"booking_hotel/helper"
	"booking_hotel/model/domain"
	"booking_hotel/model/web"
	"booking_hotel/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type MeetingRoomServiceImpl struct {
	MeetingRoomRepository repository.MeetingRoomRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewMeetingRoomService(meeting_roomRepository repository.MeetingRoomRepository, DB *sql.DB, validate *validator.Validate) MeetingRoomService {
	return &MeetingRoomServiceImpl{
		MeetingRoomRepository: meeting_roomRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *MeetingRoomServiceImpl) Create(ctx context.Context, request web.MeetingRoomCreateRequest) web.MeetingRoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meeting_room := domain.MeetingRoom{
		FloorId:    request.FloorId,
		Name:       request.Name,
		Capacity:   request.Capacity,
		FacilityId: request.FacilityId,
	}

	meeting_room = service.MeetingRoomRepository.Save(ctx, tx, meeting_room)

	return helper.ToMeetingRoomResponse(meeting_room)
}

func (service *MeetingRoomServiceImpl) Update(ctx context.Context, request web.MeetingRoomUpdateRequest) web.MeetingRoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meeting_room, err := service.MeetingRoomRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	meeting_room.FloorId = request.FloorId
	meeting_room.Name = request.Name
	meeting_room.Capacity = request.Capacity
	meeting_room.FacilityId = request.FacilityId

	meeting_room = service.MeetingRoomRepository.Update(ctx, tx, meeting_room)

	return helper.ToMeetingRoomResponse(meeting_room)
}

func (service *MeetingRoomServiceImpl) Delete(ctx context.Context, meeting_roomId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meeting_room, err := service.MeetingRoomRepository.FindById(ctx, tx, meeting_roomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MeetingRoomRepository.Delete(ctx, tx, meeting_room)
}

func (service *MeetingRoomServiceImpl) FindById(ctx context.Context, meeting_roomId int) web.MeetingRoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meeting_room, err := service.MeetingRoomRepository.FindById(ctx, tx, meeting_roomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMeetingRoomResponse(meeting_room)
}

func (service *MeetingRoomServiceImpl) FindAll(ctx context.Context) []web.MeetingRoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	meeting_rooms := service.MeetingRoomRepository.FindAll(ctx, tx)

	return helper.ToMeetingRoomResponses(meeting_rooms)
}

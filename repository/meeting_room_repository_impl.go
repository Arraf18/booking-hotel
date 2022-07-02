package repository

import (
	"booking_hotel/helper"
	"booking_hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type MeetingRoomRepositoryImpl struct {
}

func NewMeetingRoomRepository() MeetingRoomRepository {
	return &MeetingRoomRepositoryImpl{}
}

func (repository MeetingRoomRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, meeting_room domain.MeetingRoom) domain.MeetingRoom {
	SQL := "insert into meeting_rooms(floor_id, name, capacity, facility_id) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, meeting_room.FloorId, meeting_room.Name, meeting_room.Capacity, meeting_room.FacilityId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	meeting_room.Id = int(id)
	return meeting_room
}

func (repository MeetingRoomRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, meeting_room domain.MeetingRoom) domain.MeetingRoom {
	SQL := "update meeting_rooms set floor_id = ?, name = ?, capacity = ?, facility_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, meeting_room.FloorId, meeting_room.Name, meeting_room.Capacity, meeting_room.FacilityId, meeting_room.Id)
	helper.PanicIfError(err)

	return meeting_room
}

func (repository MeetingRoomRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, meeting_room domain.MeetingRoom) {
	SQL := "delete from meeting_rooms where id = ?"
	_, err := tx.ExecContext(ctx, SQL, meeting_room.Id)
	helper.PanicIfError(err)
}

func (repository MeetingRoomRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, meeting_roomId int) (domain.MeetingRoom, error) {
	SQL := "select id, floor_id, name, capacity, facility_id from meeting_rooms where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, meeting_roomId)
	helper.PanicIfError(err)
	defer rows.Close()

	meeting_room := domain.MeetingRoom{}
	if rows.Next() {
		err := rows.Scan(&meeting_room.Id, &meeting_room.FloorId, &meeting_room.Name, &meeting_room.Capacity, &meeting_room.FacilityId)
		helper.PanicIfError(err)
		return meeting_room, nil
	} else {
		return meeting_room, errors.New("Id is not found")
	}
}

func (repository MeetingRoomRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoom {
	SQL := "select id, floor_id, name, capacity, facility_id from meeting_rooms"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var meeting_rooms []domain.MeetingRoom
	for rows.Next() {
		meeting_room := domain.MeetingRoom{}
		err := rows.Scan(&meeting_room.Id, &meeting_room.FloorId, &meeting_room.Name, &meeting_room.Capacity, &meeting_room.FacilityId)
		helper.PanicIfError(err)
		meeting_rooms = append(meeting_rooms, meeting_room)
	}
	return meeting_rooms
}

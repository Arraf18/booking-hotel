package repository

import (
	"booking_hotel/model/domain"
	"context"
	"database/sql"
)

type MeetingRoomRepository interface {
	Save(ctx context.Context, tx *sql.Tx, meeting_room domain.MeetingRoom) domain.MeetingRoom
	Update(ctx context.Context, tx *sql.Tx, meeting_room domain.MeetingRoom) domain.MeetingRoom
	Delete(ctx context.Context, tx *sql.Tx, meeting_room domain.MeetingRoom)
	FindById(ctx context.Context, tx *sql.Tx, meeting_roomId int) (domain.MeetingRoom, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoom
}

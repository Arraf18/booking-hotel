package helper

import (
	"booking_hotel/model/domain"
	"booking_hotel/model/web"
)

func ToFacilityResponse(facility domain.Facility) web.FacilityResponse {
	return web.FacilityResponse{
		Id:          facility.Id,
		Name:        facility.Name,
		Description: facility.Description,
	}
}

func ToFacilityResponses(facilities []domain.Facility) []web.FacilityResponse {
	var facilityResponses []web.FacilityResponse
	for _, facility := range facilities {
		facilityResponses = append(facilityResponses, ToFacilityResponse(facility))
	}
	return facilityResponses
}
func ToMeetingRoomResponse(meeting_room domain.MeetingRoom) web.MeetingRoomResponse {
	return web.MeetingRoomResponse{
		Id:         meeting_room.Id,
		FloorId:    meeting_room.FloorId,
		Name:       meeting_room.Name,
		Capacity:   meeting_room.Capacity,
		FacilityId: meeting_room.FacilityId,
	}
}

func ToMeetingRoomResponses(meeting_rooms []domain.MeetingRoom) []web.MeetingRoomResponse {
	var meeting_roomResponses []web.MeetingRoomResponse
	for _, meeting_room := range meeting_rooms {
		meeting_roomResponses = append(meeting_roomResponses, ToMeetingRoomResponse(meeting_room))
	}
	return meeting_roomResponses
}

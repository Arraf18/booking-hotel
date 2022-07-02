package web

type MeetingRoomResponse struct {
	Id         int    `json:"id"`
	FloorId    int    `json:"floor_id"`
	Name       string `json:"name"`
	Capacity   string `json:"capacity"`
	FacilityId int    `json:"facility_id"`
}

package controller

import (
	"booking_hotel/helper"
	"booking_hotel/model/web"
	"booking_hotel/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type MeetingRoomControllerImpl struct {
	MeetingRoomService service.MeetingRoomService
}

func NewMeetingRoomController(meeting_roomService service.MeetingRoomService) MeetingRoomController {
	return &MeetingRoomControllerImpl{
		MeetingRoomService: meeting_roomService,
	}
}

func (controller *MeetingRoomControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meeting_roomCreateRequest := web.MeetingRoomCreateRequest{}
	helper.ReadFromRequestBody(request, &meeting_roomCreateRequest)

	meeting_roomResponse := controller.MeetingRoomService.Create(request.Context(), meeting_roomCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meeting_roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meeting_roomUpdateRequest := web.MeetingRoomUpdateRequest{}
	helper.ReadFromRequestBody(request, &meeting_roomUpdateRequest)

	meeting_roomId := params.ByName("meeting_roomId")
	id, err := strconv.Atoi(meeting_roomId)
	helper.PanicIfError(err)

	meeting_roomUpdateRequest.Id = id

	meeting_roomResponse := controller.MeetingRoomService.Update(request.Context(), meeting_roomUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meeting_roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meeting_roomId := params.ByName("meeting_roomId")
	id, err := strconv.Atoi(meeting_roomId)
	helper.PanicIfError(err)

	controller.MeetingRoomService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meeting_roomId := params.ByName("meeting_roomId")
	id, err := strconv.Atoi(meeting_roomId)
	helper.PanicIfError(err)

	meeting_roomResponse := controller.MeetingRoomService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meeting_roomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meeting_roomResponses := controller.MeetingRoomService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meeting_roomResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

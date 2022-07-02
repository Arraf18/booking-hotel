package app

import (
	"booking_hotel/controller"
	"booking_hotel/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(facilityController controller.FacilityController, meeting_roomController controller.MeetingRoomController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/facilities", facilityController.FindAll)
	router.GET("/api/facilities/:facilityId", facilityController.FindById)
	router.POST("/api/facilities", facilityController.Create)
	router.PUT("/api/facilities/:facilityId", facilityController.Update)
	router.DELETE("/api/facilities/:facilityId", facilityController.Delete)

	router.GET("/api/meetinf_rooms", meeting_roomController.FindAll)
	router.GET("/api/meeting_rooms/:meeting_roomId", meeting_roomController.FindById)
	router.POST("/api/meeting_rooms", meeting_roomController.Create)
	router.PUT("/api/meeting_rooms/:meeting_roomId", meeting_roomController.Update)
	router.DELETE("/api/meeitnf_rooms/:meeting_roomId", meeting_roomController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

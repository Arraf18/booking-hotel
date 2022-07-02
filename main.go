package main

import (
	"booking_hotel/app"
	"booking_hotel/controller"
	"booking_hotel/helper"
	"booking_hotel/middleware"
	"booking_hotel/repository"
	"booking_hotel/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	facilityRepository := repository.NewFacilityRepository()
	facilityService := service.NewFacilityService(facilityRepository, db, validate)
	facilityController := controller.NewFacilityController(facilityService)

	meeting_roomRepository := repository.NewMeetingRoomRepository()
	meeting_roomService := service.NewMeetingRoomService(meeting_roomRepository, db, validate)
	meeting_roomController := controller.NewMeetingRoomController(meeting_roomService)

	router := app.NewRouter(facilityController, meeting_roomController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

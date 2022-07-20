// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"playground.io/another-pet-store/controller"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/service"
)

// Injectors from controller_injector.go:

func NewAnimalController() controller.AnimalController {
	connection := db.GetConnection()
	profileRepository := db.NewProfileRepository(connection)
	profileService := service.NewProfileService(profileRepository)
	animalRepository := db.NewAnimalRepository(connection)
	animalService := service.NewAnimalService(profileService, animalRepository)
	animalController := controller.NewAnimalController(animalService)
	return animalController
}

func NewLoginController() controller.LoginController {
	connection := db.GetConnection()
	userRepository := db.NewUserRepository(connection)
	userService := service.NewUserService(userRepository)
	profileRepository := db.NewProfileRepository(connection)
	profileService := service.NewProfileService(profileRepository)
	jwtService := service.NewJWTService()
	loginService := service.NewLoginService(userService, profileService, jwtService)
	loginController := controller.NewLoginController(loginService)
	return loginController
}

func NewProfileController() controller.ProfileController {
	connection := db.GetConnection()
	profileRepository := db.NewProfileRepository(connection)
	profileService := service.NewProfileService(profileRepository)
	profileController := controller.NewProfileController(profileService)
	return profileController
}

func NewReferenceController() controller.ReferenceController {
	connection := db.GetConnection()
	referenceRepository := db.NewReferenceRepository(connection)
	referenceService := service.NewReferenceService(referenceRepository)
	referenceController := controller.NewReferenceController(referenceService)
	return referenceController
}

func NewSpecialOfferController() controller.SpecialOfferController {
	connection := db.GetConnection()
	specialOfferRepository := db.NewSpecialOfferRepository(connection)
	specialOfferService := service.NewSpecialOfferService(specialOfferRepository)
	specialOfferController := controller.NewSpecialOfferController(specialOfferService)
	return specialOfferController
}

func NewChatController() controller.ChatController {
	connection := db.GetConnection()
	chatRepository := db.NewChatRepository(connection)
	chatService := service.NewChatService(chatRepository)
	ticketService := service.NewTicketService()
	chatController := controller.NewChatController(chatService, ticketService, connection)
	return chatController
}

func NewStickerController() controller.StickerController {
	connection := db.GetConnection()
	stickerRepository := db.NewStickerRepository(connection)
	stickerService := service.NewStickerService(stickerRepository)
	stickerController := controller.NewStickerController(stickerService)
	return stickerController
}

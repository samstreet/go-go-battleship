package services

import (
	BoardModels "../../board/model"
	"../../core/dbal"
	"../model"
)

type SessionService struct {
	Model model.SessionModel
}

func NewSessionService() *SessionService {
	return &SessionService{
		Model: *model.NewSessionModel(),
	}
}

func (service SessionService) CreateSession() model.SessionModel {

	board := BoardModels.BoardModel{}
	dbal.InitialiseConnection().Create(&board)

	session := model.SessionModel{Board:board}
	dbal.InitialiseConnection().Create(&session)

	dbal.InitialiseConnection().Model(&session).Related(&board)

	return session
}

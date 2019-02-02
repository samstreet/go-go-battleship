package services

import (
	BoardModels "../../board/model"
	BoardDTO "../../board/structs"
	"../../core/dbal"
	"../model"
	"../structs"
	u "github.com/satori/go.uuid"
)

type SessionService struct {
	Model model.SessionModel
}

func NewSessionService() *SessionService {
	m := model.SessionModel{}

	return &SessionService{
		Model: *m.Fresh(),
	}
}

func (service SessionService) CreateSession() *model.SessionModel {
	db := dbal.InitialiseConnection()

	board := BoardModels.BoardModel{}
	db.Create(&board)

	session := service.Model.Fresh()
	session.Board = board

	db.Create(&session)

	return session
}

func (service SessionService) FindSessionByUUID(uuid string) structs.SessionOutDTO {
	sessionOutDTO := structs.SessionOutDTO{}
	boardOutDTO := BoardDTO.BoardOutDTO{}
	tmp := service.Model.Fresh()

	service.Model.Connection.Where("id = ?", uuid).First(&tmp)

	boardOutDTO.UUID = u.FromStringOrNil(tmp.BoardID)

	sessionOutDTO.UUID = u.FromStringOrNil(tmp.ID)
	sessionOutDTO.Board = boardOutDTO
	return sessionOutDTO
}

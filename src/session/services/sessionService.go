package services

import (
	BoardModels "../../board/model"
	BoardDTO "../../board/structs"
	"../../core/dbal"
	"../model"
	SessionDTO "../structs"
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

func (service SessionService) CreateSession() SessionDTO.SessionOutDTO {
	db := dbal.InitialiseConnection()

	board := BoardModels.BoardModel{}
	db.Create(&board)

	session := service.Model.Fresh()
	session.Board = board

	db.Create(&session)

	sessionOutDTO := SessionDTO.SessionOutDTO{}
	boardDTO := BoardDTO.BoardOutDTO{UUID: u.FromStringOrNil(session.Board.ID)}

	sessionOutDTO.Board = boardDTO
	sessionOutDTO.UUID = u.FromStringOrNil(session.ID)

	return sessionOutDTO
}

func (service SessionService) FindSessionByUUID(uuid string) SessionDTO.SessionOutDTO {
	sessionOutDTO := SessionDTO.SessionOutDTO{}
	boardOutDTO := BoardDTO.BoardOutDTO{}
	tmp := service.Model.Fresh()

	service.Model.Connection.Where("id = ?", uuid).First(&tmp)

	boardOutDTO.UUID = u.FromStringOrNil(tmp.BoardID)

	sessionOutDTO.UUID = u.FromStringOrNil(tmp.ID)
	sessionOutDTO.Board = boardOutDTO
	return sessionOutDTO
}

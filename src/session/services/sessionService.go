package services

import (
	BoardModels "../../board/model"
	CoreModels "../../core/model"
	BoardDTO "../../board/structs"
	CoreDTO "../../core/structs"
	"../../core/dbal"
	CoreServices "../../core/services"
	"../model"
	SessionDTO "../structs"
	u "github.com/satori/go.uuid"
)

type SessionService struct {
	Model       model.SessionModel
	UserService CoreServices.UserService
}

func NewSessionService() *SessionService {
	m := model.SessionModel{}

	return &SessionService{
		Model:       *m.Fresh(),
		UserService: *CoreServices.NewUserService(),
	}
}

func (service SessionService) CreateSession(dto SessionDTO.CreateSessionDTO) SessionDTO.SessionOutDTO {
	db := dbal.InitialiseConnection()

	board := BoardModels.BoardModel{}
	board.XLength = 10
	board.YLength = board.XLength

	user := dto.Player1

	session := model.SessionModel{Board: board, Players: []CoreModels.User{user}}
	db.Save(&session)

	sessionOutDTO := SessionDTO.SessionOutDTO{}
	boardOutDTO := BoardDTO.BoardOutDTO{UUID: u.FromStringOrNil(session.Board.ID), Players:[]CoreDTO.UserOutDTO{
		{UUID:user.GetID()},
	},}
	boardOutDTO.XLength = board.XLength
	boardOutDTO.YLength = board.YLength

	sessionOutDTO.Board = boardOutDTO
	sessionOutDTO.UUID = u.FromStringOrNil(session.ID)

	return sessionOutDTO
}

func (service SessionService) FindSessionByUUID(uuid string) SessionDTO.SessionOutDTO {
	boardOutDTO := BoardDTO.BoardOutDTO{}
	tmp := service.Model.Fresh()

	service.Model.Connection.Preload("Board").Where("id = ?", uuid).First(&tmp)

	boardOutDTO.UUID = u.FromStringOrNil(tmp.Board.ID)
	boardOutDTO.XLength = tmp.Board.XLength
	boardOutDTO.YLength = tmp.Board.YLength

	return SessionDTO.SessionOutDTO{UUID: u.FromStringOrNil(tmp.ID), Board: boardOutDTO}
}

func (service SessionService) JoinSession(request SessionDTO.JoinSessionDTO) SessionDTO.JoinSessionDTO {
	return request
}

func (service SessionService) attachUserToSession(session u.UUID, playerKey u.UUID) bool {
	//user := service.UserService.findByUUID(playerKey)
	return true
}

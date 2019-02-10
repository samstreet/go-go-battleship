package services

import (
	BoardModels "../../board/model"
	BoardDTO "../../board/structs"
	"../../core/dbal"
	CoreModels "../../core/model"
	CoreServices "../../core/services"
	CoreDTO "../../core/structs"
	"../model"
	SessionDTO "../structs"
	u "github.com/satori/go.uuid"
	"log"
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
	user.PlayerRole = 1
	db.Save(&user)

	session := model.SessionModel{Board: board, Players: []CoreModels.User{user}}
	db.Save(&session)

	sessionOutDTO := SessionDTO.SessionOutDTO{}
	boardOutDTO := BoardDTO.BoardOutDTO{UUID: u.FromStringOrNil(session.Board.ID), Players: []CoreDTO.UserOutDTO{
		{UUID: user.GetID(), PlayerRole: user.PlayerRole},
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

	service.Model.Connection.Preload("Board").Preload("Players").Where("id = ?", uuid).First(&tmp)

	boardOutDTO.UUID = u.FromStringOrNil(tmp.Board.ID)
	boardOutDTO.XLength = tmp.Board.XLength
	boardOutDTO.YLength = tmp.Board.YLength

	log.Println(len(tmp.Players))
	var lol []CoreDTO.UserOutDTO

	for i := 0; i < len(tmp.Players); i++ {
		log.Println(i)
		lol = append(lol, CoreDTO.UserOutDTO{UUID: tmp.Players[i].GetID(), PlayerRole: tmp.Players[i].PlayerRole})
	}

	boardOutDTO.Players = lol

	return SessionDTO.SessionOutDTO{
		UUID:  u.FromStringOrNil(tmp.ID),
		Board: boardOutDTO,
	}
}

func (service SessionService) JoinSession(request SessionDTO.JoinSessionDTO) bool {
	user, _ := CoreServices.NewUserService().FindByUUID(request.Player2.String())
	user.PlayerRole = 2
	user.SessionID = request.UUID.String()

	db := dbal.InitialiseConnection()
	db.Save(&user)

	// todo - check if there are already two players attached to a game

	return true
}

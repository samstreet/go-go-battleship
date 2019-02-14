package session

import (
	. "../../board/model"
	. "../../board/structs"
	. "../../core/dbal"
	. "../../core/model"
	. "../../core/services"
	. "../../core/structs"
	. "../model"
	. "../structs"
	u "github.com/satori/go.uuid"
	"log"
)

type SessionService struct {
	Model       SessionModel
	UserService UserService
}

func NewSessionService() *SessionService {
	m := SessionModel{}

	return &SessionService{
		Model:       *m.Fresh(),
		UserService: *NewUserService(),
	}
}

func (service SessionService) CreateSession(dto CreateSessionDTO) SessionOutDTO {
	db := InitialiseConnection()

	board := BoardModel{}
	board.XLength = 10
	board.YLength = board.XLength

	user := dto.Player1
	user.PlayerRole = 1
	db.Save(&user)

	session := SessionModel{Board: board, Players: []User{user}}
	db.Save(&session)

	sessionOutDTO := SessionOutDTO{}
	boardOutDTO := BoardOutDTO{UUID: u.FromStringOrNil(session.Board.ID), Players: []UserOutDTO{
		{UUID: user.GetID(), PlayerRole: user.PlayerRole},
	},}
	boardOutDTO.XLength = board.XLength
	boardOutDTO.YLength = board.YLength

	sessionOutDTO.Board = boardOutDTO
	sessionOutDTO.UUID = u.FromStringOrNil(session.ID)

	return sessionOutDTO
}

func (service SessionService) FindSessionByUUID(uuid string) SessionOutDTO {
	boardOutDTO := BoardOutDTO{}
	tmp := service.Model.Fresh()

	service.Model.Connection.Preload("Board").Preload("Players").Where("id = ?", uuid).First(&tmp)

	boardOutDTO.UUID = u.FromStringOrNil(tmp.Board.ID)
	boardOutDTO.XLength = tmp.Board.XLength
	boardOutDTO.YLength = tmp.Board.YLength

	log.Println(len(tmp.Players))
	var lol []UserOutDTO

	for i := 0; i < len(tmp.Players); i++ {
		log.Println(i)
		lol = append(lol, UserOutDTO{UUID: tmp.Players[i].GetID(), PlayerRole: tmp.Players[i].PlayerRole})
	}

	boardOutDTO.Players = lol

	return SessionOutDTO{
		UUID:  u.FromStringOrNil(tmp.ID),
		Board: boardOutDTO,
	}
}

func (service SessionService) JoinSession(request JoinSessionDTO) bool {
	user, _ := NewUserService().FindByUUID(request.Player2.String())
	user.PlayerRole = 2
	user.SessionID = request.UUID.String()

	db := InitialiseConnection()
	db.Save(&user)

	// todo - check if there are already two players attached to a game

	return true
}

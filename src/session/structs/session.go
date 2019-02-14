package session

import (
	. "../../board/structs"
	. "../../core/model"
	"encoding/xml"
	"github.com/satori/go.uuid"
)

type SessionOutDTO struct {
	XMLName xml.Name    `xml:"Session" json:"-"`
	UUID    uuid.UUID   `xml:"UUID" json:"uuid"`
	Board   BoardOutDTO `xml:"Board" json:"board"`
}

type CreateSessionDTO struct {
	Player1 User `xml:"Player1" json:"player1>uuid"`
}

type JoinSessionDTO struct {
	XMLName xml.Name  `xml:"Session" json:"-"`
	UUID    uuid.UUID `xml:"UUID" json:"uuid"`
	Player1 uuid.UUID `xml:"Player1>UUID" json:"player1>uuid"`
	Player2 uuid.UUID `xml:"Player2>UUID" json:"player2>uuid"`
}

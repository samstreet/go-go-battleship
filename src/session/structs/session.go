package structs

import (
	"../../core/model"
	"encoding/xml"
	"github.com/satori/go.uuid"
)
import "../../board/structs"

type SessionOutDTO struct {
	XMLName xml.Name            `xml:"Session" json:"-"`
	UUID    uuid.UUID           `xml:"UUID" json:"uuid"`
	Board   structs.BoardOutDTO `xml:"Board" json:"board"`
}

type CreateSessionDTO struct {
	Player1 model.User `xml:"Player1" json:"player1>uuid"`
}

type JoinSessionDTO struct {
	XMLName xml.Name  `xml:"Session" json:"-"`
	UUID    uuid.UUID `xml:"UUID" json:"uuid"`
	Player1 uuid.UUID `xml:"Player1>UUID" json:"player1>uuid"`
	Player2 uuid.UUID `xml:"Player2>UUID" json:"player2>uuid"`
}

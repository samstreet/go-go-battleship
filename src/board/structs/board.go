package board

import (
	. "../../core/structs"
	"encoding/xml"
	"github.com/satori/go.uuid"
)

type BoardOutDTO struct {
	XMLName xml.Name             `xml:"Board" json:"-"`
	UUID    uuid.UUID            `xml:"UUID" json:"uuid"`
	XLength int                  `xml:"XLength,attr" json:"xLength"`
	YLength int                  `xml:"YLength,attr" json:"yLength"`
	Players []UserOutDTO `xml:"Players>Player" json:"players"`
}

package structs

import (
	"../../core/structs"
	"encoding/xml"
	"github.com/satori/go.uuid"
)

type BoardOutDTO struct {
	XMLName xml.Name             `xml:"Board" json:"-"`
	UUID    uuid.UUID            `xml:"UUID" json:"uuid"`
	XLength int                  `xml:"XLength" json:"xLength"`
	YLength int                  `xml:"YLength" json:"yLength"`
	Players []structs.UserOutDTO `xml:"Players>Player" json:"players"`
}

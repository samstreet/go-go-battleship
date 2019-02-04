package structs

import (
	"encoding/xml"
	"github.com/satori/go.uuid"
)
import "../../board/structs"

type SessionOutDTO struct {
	XMLName xml.Name `xml:"Session" json:"-"`
	UUID uuid.UUID `xml:"UUID"`
	Board structs.BoardOutDTO `xml:"Board"`
}

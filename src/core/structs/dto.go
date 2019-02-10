package structs

import "github.com/satori/go.uuid"

type UserOutDTO struct {
	UUID uuid.UUID
	PlayerRole int `xml:"player_role,attr" json:"player_role"`
}

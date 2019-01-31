package structs

import "github.com/satori/go.uuid"

type SessionOutDTO struct {
	UUID uuid.UUID
	Board uuid.UUID
}

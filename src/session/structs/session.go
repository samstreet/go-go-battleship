package structs

import "github.com/satori/go.uuid"
import "../../board/structs"

type SessionOutDTO struct {
	UUID uuid.UUID
	Board structs.BoardOutDTO
}

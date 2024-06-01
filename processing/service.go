package processing

import "github.com/google/uuid"

func StartProcessing() uuid.UUID {
	id, _ := uuid.NewRandom()

	return id
}

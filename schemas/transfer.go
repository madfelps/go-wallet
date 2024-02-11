package schemas

import "github.com/google/uuid"

type Transfer struct {
	Amount int       `json:"amount"`
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

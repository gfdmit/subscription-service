package model

import (
	"github.com/gfdmit/subscription-service/internal/utils"
	"github.com/google/uuid"
)

type Subscription struct {
	ID          int               `json:"id"`
	ServiceName string            `json:"service_name"`
	Price       int               `json:"price"`
	UserID      uuid.UUID         `json:"user_id"`
	StartDate   utils.CustomDate  `json:"start_date"`
	EndDate     *utils.CustomDate `json:"end_date,omitempty"`
}

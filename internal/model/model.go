package model

import (
	"github.com/gfdmit/subscription-service/internal/utils"
)

type Subscription struct {
	ID          string            `json:"id"`
	ServiceName string            `json:"service_name"`
	Price       int               `json:"price"`
	UserID      string            `json:"user_id"`
	StartDate   utils.CustomDate  `json:"start_date"`
	EndDate     *utils.CustomDate `json:"end_date,omitempty"`
}

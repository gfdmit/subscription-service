package repository

import (
	"context"

	"github.com/gfdmit/subscription-service/internal/model"
)

type Repository interface {
	CreateSubscription(ctx context.Context, subscription model.Subscription) (*model.Subscription, error)
	GetSubscription(ctx context.Context, id int) (*model.Subscription, error)
	GetSubscriptions(ctx context.Context) ([]model.Subscription, error)
	UpdateSubscription(ctx context.Context, id int, subscription model.Subscription) (*model.Subscription, error)
	DeleteSubscription(ctx context.Context, id int) (*model.Subscription, error)
	GetAmount(ctx context.Context, activeParams map[string]string) (int, error)
}

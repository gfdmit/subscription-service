package subscription

import (
	"context"
	"errors"

	"github.com/gfdmit/subscription-service/internal/model"
	"github.com/gfdmit/subscription-service/internal/repository"
)

type subscriptionService struct {
	repo repository.Repository
}

func New(repo repository.Repository) *subscriptionService {
	return &subscriptionService{repo: repo}
}

func (ss subscriptionService) CreateSubscription(ctx context.Context, subscription model.Subscription) (*model.Subscription, error) {
	if subscription.Price < 0 {
		return nil, errors.New("price is negative")
	}
	return ss.repo.CreateSubscription(ctx, subscription)
}

func (ss subscriptionService) GetSubscription(ctx context.Context, id int) (*model.Subscription, error) {
	return ss.repo.GetSubscription(ctx, id)
}

func (ss subscriptionService) GetSubscriptions(ctx context.Context) ([]model.Subscription, error) {
	return ss.repo.GetSubscriptions(ctx)
}

func (ss subscriptionService) UpdateSubscription(ctx context.Context, id int, subscription model.Subscription) (*model.Subscription, error) {
	return ss.repo.UpdateSubscription(ctx, id, subscription)
}

func (ss subscriptionService) DeleteSubscription(ctx context.Context, id int) (*model.Subscription, error) {
	return ss.repo.DeleteSubscription(ctx, id)
}

func (ss subscriptionService) GetAmount(ctx context.Context, activeParams map[string]string) (int, error) {
	return ss.repo.GetAmount(ctx, activeParams)
}

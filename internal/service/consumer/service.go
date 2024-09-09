package consumerService

import "context"

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}
package user_saver

import (
	"context"

	"github.com/lookandhate/course_auth/internal/config"
	"github.com/lookandhate/course_auth/internal/repository"
	"github.com/lookandhate/course_auth/internal/service/consumer"
	kafka "github.com/lookandhate/course_platform_lib/pkg/message_queue/kafka/client"
)

var _ consumerService.ConsumerService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	consumer       kafka.Consumer
	config         config.KafkaConfig
}

// NewService creates a new user saver service.
func NewService(
	userRepository repository.UserRepository,
	consumer kafka.Consumer,
	kafkaConfig config.KafkaConfig,
) *service {
	return &service{
		userRepository: userRepository,
		consumer:       consumer,
		config:         kafkaConfig,
	}
}

// RunConsumer runs user saver consumer.
func (s *service) RunConsumer(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-s.run(ctx):
			if err != nil {
				return err
			}
		}
	}
}

func (s *service) run(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		defer close(errChan)

		errChan <- s.consumer.Consume(ctx, s.config.TopicName, s.UserSaveHandler)
	}()

	return errChan
}

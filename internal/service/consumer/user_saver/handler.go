package user_saver

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/lookandhate/course_auth/internal/service/model"
)

func (s *service) UserSaveHandler(ctx context.Context, msg *sarama.ConsumerMessage) error {
	userInfo := &model.CreateUserModel{}
	err := json.Unmarshal(msg.Value, userInfo)
	if err != nil {
		return err
	}

	_, err = s.userRepository.CreateUser(ctx, userInfo)
	if err != nil {
		return err
	}

	return nil
}

package user_saver

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	serviceLayer "github.com/lookandhate/course_auth/internal/service"
	"github.com/lookandhate/course_auth/internal/service/model"
)

func (s *service) UserSaveHandler(ctx context.Context, msg *sarama.ConsumerMessage) error {
	user := &model.CreateUserModel{}
	err := json.Unmarshal(msg.Value, user)
	if err != nil {
		return err
	}

	// Check user role has been passed correctly
	if user.Role == model.UserUnknownRole {
		return serviceLayer.ErrInvalidRole
	}

	if user.PasswordConfirm != user.Password {
		return serviceLayer.ErrPasswordMismatch
	}

	_, err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"context"

	"github.com/lookandhate/course_auth/internal/service"
)

// validateID check passed ID and returns error if it is not correct.
func (s *Service) validateID(id int) error {
	if id <= 0 {
		return service.ErrInvalidID
	}
	return nil
}

// checkUserExists check if user exists using repo method
// returns service error if user does not exist or repo could not make a request.
func (s *Service) checkUserExists(ctx context.Context, id int) error {
	isExists, err := s.repo.CheckUserExists(ctx, id)
	if err != nil {
		// TODO handle as internal server error
		return err
	}
	if !isExists {
		return service.ErrUserDoesNotExist
	}
	return nil
}

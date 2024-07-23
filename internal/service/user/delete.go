package service

import (
	"context"
)

// DeleteUser deletes user by given ID if it is correct.
func (s *Service) Delete(ctx context.Context, id int) error {
	if err := s.validateID(id); err != nil {
		return err
	}

	return s.repo.DeleteUser(ctx, id)
}

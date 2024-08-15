package service

import (
	"context"
	"log"
)

// Delete deletes user by given ID if it is correct.
func (s *Service) Delete(ctx context.Context, id int) error {
	if err := s.validateID(id); err != nil {
		return err
	}

	err := s.cache.Delete(ctx, id)
	if err != nil {
		log.Default().Println(err)
	}

	return s.repo.DeleteUser(ctx, id)
}

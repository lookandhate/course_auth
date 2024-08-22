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
		log.Printf("Error when deleting user from cache: %v", err)
	}

	err = s.repo.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("Error when deleting user from repo: %v", err)
		return err
	}

	return nil
}

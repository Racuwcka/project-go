package item

import "context"

type DeleteRepository interface {
	DeleteItem(int64, uint32) error
}

type DeleteService struct {
	name string
	repo DeleteRepository
}

func NewDeleteService(repo DeleteRepository) *DeleteService {
	return &DeleteService{
		name: "item delete service",
		repo: repo,
	}
}

func (s DeleteService) Delete(_ context.Context, user int64, sku uint32) error {
	if err := s.repo.DeleteItem(user, sku); err != nil {
		return err
	}
	return nil
}

package item

type ClearRepository interface {
	DeleteItemsByUserID(int64) error
}

type ClearService struct {
	name string
	repo ClearRepository
}

func NewClearService(repo ClearRepository) *ClearService {
	return &ClearService{
		name: "item clear service",
		repo: repo,
	}
}

func (s ClearService) Clear(user int64) error {
	if err := s.repo.DeleteItemsByUserID(user); err != nil {
		return err
	}
	return nil
}

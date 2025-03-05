package item

type CheckoutRepository interface {
	GetItemsByUserID(int64) (map[uint32]uint16, error)
	DeleteItemsByUserID(int64) error
}

type CheckoutService struct {
	name string
	repo CheckoutRepository
}

func NewCheckoutService(repo CheckoutRepository) *CheckoutService {
	return &CheckoutService{
		name: "item checkout service",
		repo: repo,
	}
}

func (s CheckoutService) Checkout(user int64) error {
	_, err := s.repo.GetItemsByUserID(user)
	if err != nil {
		return err
	}

	// order/create

	if err := s.repo.DeleteItemsByUserID(user); err != nil {
		return err
	}
	return nil
}

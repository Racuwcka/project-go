package item

import "context"

type DeleteService struct {
	name string
}

func NewDeleteService() *DeleteService {
	return &DeleteService{
		name: "item delete service",
	}
}

func (s DeleteService) Delete(ctx context.Context, user int64, sku uint32) error {
	return nil
}

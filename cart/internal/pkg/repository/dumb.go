package repository

import "context"

type DumbRepo struct{}

func NewDumbRepo() *DumbRepo {
	return &DumbRepo{}
}

func (d DumbRepo) AddItem(_ context.Context, _ int64, _ uint32, _ uint16) error {
	return nil
}

func (d DumbRepo) GetItemsByUserID(_ int64) (map[uint32]uint16, error) {
	return map[uint32]uint16{3618852: 2, 4288068: 3, 4487693: 1}, nil
}

func (d DumbRepo) DeleteItem(_ int64, _ uint32) error {
	return nil
}

func (d DumbRepo) DeleteItemsByUserID(_ int64) error {
	return nil
}

package repository

import "math/rand"

type DumbRepo struct{}

func NewDumbRepo() *DumbRepo {
	return &DumbRepo{}
}

func (DumbRepo) GetStocks(_ uint32) uint64 {
	return uint64(rand.Int() % 1000)
}

package item

import (
	"context"
	"errors"
	"fmt"
	"route256/cart/internal/pkg/clients/product"
	"time"
)

type AddRepository interface {
	AddItem(ctx context.Context, user int64, sku uint32, count uint16) error
}

type StocksProvider interface {
	GetStocks(ctx context.Context, sku uint32) (uint64, error)
}

type AddService struct {
	name            string
	stocksProvider  StocksProvider
	productProvider product.ProductProvider
	repo            AddRepository
}

var ErrInsufficientStocks = errors.New("insufficient stocks")

func NewAddService(stocksProvider StocksProvider, productProvider product.ProductProvider) *AddService {
	return &AddService{
		name:            "item add service",
		stocksProvider:  stocksProvider,
		productProvider: productProvider,
	}
}

func (s AddService) Add(ctx context.Context, user int64, sku uint32, count uint16, token string) error {
	if _, _, err := s.productProvider.GetProductInfo(sku, token); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	stocksCount, err := s.stocksProvider.GetStocks(ctx, sku)
	if err != nil {
		return err
	}

	if uint64(count) > stocksCount {
		return fmt.Errorf("%w: %d stocks available", ErrInsufficientStocks, count)
	}

	if err := s.repo.AddItem(ctx, user, sku, count); err != nil {
		return err
	}

	return nil
}

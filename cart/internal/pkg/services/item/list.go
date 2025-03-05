package item

import (
	"fmt"
	"route256/cart/internal/pkg/clients/product"
	"route256/cart/internal/pkg/handlers/item"
)

type GetRepository interface {
	GetItemsByUserID(int64) (map[uint32]uint16, error)
}

type ListService struct {
	name            string
	productProvider product.ProductProvider
	repo            GetRepository
}

func NewListService(productProvider product.ProductProvider, repo GetRepository) *ListService {
	return &ListService{
		name:            "item list service",
		productProvider: productProvider,
		repo:            repo,
	}
}

func (l ListService) List(user int64, token string) ([]item.Item, uint32, error) {
	skuItems, err := l.repo.GetItemsByUserID(user)
	if err != nil {
		return []item.Item{}, 0, fmt.Errorf("failed to get items from repo: %w", err)
	}

	res := make([]item.Item, 0, len(skuItems))
	var totalPrice uint32

	for sku, count := range skuItems {
		name, price, err := l.productProvider.GetProductInfo(sku, token)
		if err != nil {
			return []item.Item{}, 0, fmt.Errorf("failed to get product info for sku %d: %w", sku, err)
		}

		res = append(res, item.Item{
			SKU:   sku,
			Count: count,
			Name:  name,
			Price: price,
		})

		totalPrice += price * uint32(count)
	}

	return res, totalPrice, nil
}

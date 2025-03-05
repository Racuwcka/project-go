package services

type InfoRepository interface {
	GetBySKU(sku uint32) uint64
}

type StocksService struct {
	repo InfoRepository
}

func NewStocksService(repo InfoRepository) *StocksService {
	return &StocksService{
		repo: repo,
	}
}

func (s *StocksService) GetStocks(sku uint32) uint64 {
	count := s.repo.GetBySKU(sku)
	return count
}

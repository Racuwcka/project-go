package item

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg/handlers"
)

func (r ListRequest) Validate() error {
	if r.User == 0 {
		return errorIncorrectUser
	}
	return nil
}

type ListRequest struct {
	User int64 `json:"user,omitempty"`
}

type Item struct {
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
	Name  string `json:"name,omitempty"`
	Price uint32 `json:"price,omitempty"`
}

type ListResponse struct {
	Items      []Item `json:"items,omitempty"`
	TotalPrice uint32 `json:"totalPrice,omitempty"`
}

type Lister interface {
	List(user int64, token string) ([]Item, uint32, error)
}

type ListHandler struct {
	name string
	l    Lister
}

func NewItemsListHandler(lister Lister) *ListHandler {
	return &ListHandler{
		name: "item list handler",
		l:    lister,
	}
}

func (h ListHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		handlers.GetErrorResponse(w, h.name, errorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	req := &ListRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	items, totalPrice, err := h.l.List(req.User, token)
	if err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	listResponse := ListResponse{
		Items:      items,
		TotalPrice: totalPrice,
	}

	raw, err := json.Marshal(listResponse)
	if err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	handlers.GetSuccessResponse(w, raw)
}

package item

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg/handlers"
)

func (r AddRequest) Validate() error {
	if r.User == 0 {
		return errorIncorrectUser
	}
	if r.SKU == 0 {
		return errorIncorrectSKU
	}
	if r.Count == 0 {
		return errorIncorrectCount
	}
	return nil
}

type AddRequest struct {
	User  int64  `json:"user,omitempty"`
	SKU   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

type Adder interface {
	Add(ctx context.Context, user int64, sku uint32, count uint16, token string) error
}

type AddHandler struct {
	name string
	s    Adder
}

func NewItemsAddHandler(itemsAdder Adder) *AddHandler {
	return &AddHandler{
		name: "item add handler",
		s:    itemsAdder,
	}
}

func (h AddHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	req := &AddRequest{}
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

	if err := h.s.Add(r.Context(), req.User, req.SKU, req.Count, token); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

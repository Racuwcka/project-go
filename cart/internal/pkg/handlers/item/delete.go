package item

import (
	"context"
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg/handlers"
)

type DeleteHandler struct {
	name string
	d    Deleter
}

type DeleteRequest struct {
	User int64  `json:"user,omitempty"`
	Sku  uint32 `json:"sku,omitempty"`
}

func (r DeleteRequest) Validate() error {
	if r.User == 0 {
		return errorIncorrectUser
	}
	if r.Sku == 0 {
		return errorIncorrectSKU
	}
	return nil
}

type DeleteResponse struct{}

type Deleter interface {
	Delete(ctx context.Context, user int64, sku uint32) error
}

func NewItemsDeleteHandler(deleter Deleter) *DeleteHandler {
	return &DeleteHandler{
		name: "item delete handler",
		d:    deleter,
	}
}

func (h DeleteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	req := &DeleteRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
	}

	if err := h.d.Delete(r.Context(), req.User, req.Sku); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

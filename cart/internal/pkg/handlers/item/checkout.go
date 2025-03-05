package item

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg/handlers"
)

type CheckoutService interface {
	Checkout(int64) error
}

type CheckoutHandler struct {
	name string
	c    CheckoutService
}

func NewItemsCheckoutHandler(c CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{
		name: "item checkout handler",
		c:    c,
	}
}

type CheckoutRequest struct {
	User int64 `json:"user,omitempty"`
}

type CheckoutResponse struct {
	OrderID int64 `json:"order_id,omitempty"`
}

func (r CheckoutRequest) Validate() error {
	if r.User == 0 {
		return errorIncorrectUser
	}
	return nil
}

func (h CheckoutHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		handlers.GetErrorResponse(w, h.name, errorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	req := &CheckoutRequest{}
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

	if err := h.c.Checkout(req.User); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
	}

	//raw, err := json.Marshal(listResponse)
	//if err != nil {
	//	handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//handlers.GetSuccessResponse(w, raw)

	w.WriteHeader(http.StatusOK)
}

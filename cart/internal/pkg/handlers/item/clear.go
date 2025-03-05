package item

import (
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg/handlers"
)

type Clearer interface {
	Clear(int64) error
}

type ClearHandler struct {
	name string
	c    Clearer
}

func NewItemsClearHandler(c Clearer) *ClearHandler {
	return &ClearHandler{
		name: "item clear handler",
		c:    c,
	}
}

type ClearRequest struct {
	User int64 `json:"user,omitempty"`
}

func (c ClearRequest) Validate() error {
	if c.User == 0 {
		return errorIncorrectUser
	}
	return nil
}

func (h ClearHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	req := &ClearRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err := h.c.Clear(req.User); err != nil {
		handlers.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

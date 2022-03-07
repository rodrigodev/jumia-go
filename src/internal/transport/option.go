package transport

import (
	"github.com/rodrigodev/jumia-go/src/internal/phone/service"
	"net/http"
)

// Option is a Handler modifier.
type Option func(*Handler) error

// Health endpoints
func Health() Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/_live", func(w http.ResponseWriter, r *http.Request) {})
		return nil
	}
}

// Static endpoints
func Static() Option {
	return func(h *Handler) error {
		fileServer := http.FileServer(http.Dir("./static/"))
		h.router.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))
		return nil
	}
}

// Phone endpoints
func Phone(s service.PhoneServiceContainer) Option {
	return func(h *Handler) error {
		h.router.HandleFunc("/phone", GetPhonesHandler(s)).Methods("GET", "OPTIONS")
		return nil
	}
}

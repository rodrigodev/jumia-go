package transport

import (
	"net/http"

	"github.com/rodrigodev/jumia-go/src/internal/infrastructure"
	"github.com/rodrigodev/jumia-go/src/internal/phone/service"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	countryParameter = "country"
	stateParameter   = "state"
)

const errGetPhonesFailed = "fetching phones was not possible"

type Handler struct{ router *mux.Router }

func New(opts ...Option) (*Handler, error) {
	h := &Handler{router: mux.NewRouter()}

	if err := h.Apply(opts...); err != nil {
		return nil, err
	}

	return h, nil
}

// Apply will apply the options passed.
func (h *Handler) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(h); err != nil {
			return err
		}
	}
	return nil
}

// ServeHTTP calls the serve method on the router
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// GetPhonesHandler gets all phones
func GetPhonesHandler(s *service.PhoneService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		phones, err := s.GetPhones(ctx)
		if err != nil {
			infrastructure.Logger(ctx).Error(errGetPhonesFailed, zap.Error(err))
			WriteError(w, errGetPhonesFailed, http.StatusUnauthorized)
		}

		WriteJson(w, http.StatusOK, phones)
	}
}

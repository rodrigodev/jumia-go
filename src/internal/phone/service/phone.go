//go:generate go run -mod=mod github.com/golang/mock/mockgen -package service -destination=./phone_service_mock.go -source=./phone.go -build_flags=-mod=mod
package service

import (
	"context"
	"github.com/rodrigodev/jumia-go/src/internal/phone/model"
	"github.com/rodrigodev/jumia-go/src/internal/phone/repository"
	"github.com/rodrigodev/jumia-go/src/internal/phone/value"

	_ "github.com/golang/mock/mockgen/model"

	"regexp"
	"strings"
)

// PhoneServiceContainer describes the service.
type PhoneServiceContainer interface {
	GetPhones(ctx context.Context) ([]value.Phone, error)
}

// PhoneService @impl Service static check
type PhoneService struct {
	repository repository.CustomerManager
}

// NewPhoneService responsible for doing business logic for phone
func NewPhoneService(repository repository.CustomerManager) *PhoneService {
	return &PhoneService{
		repository,
	}
}

// GetPhones fetch all phones and build the data for display
func (s *PhoneService) GetPhones(ctx context.Context) ([]value.Phone, error) {
	customers, err := s.repository.GetAllCustomer(ctx)
	if err != nil {
		return nil, err
	}

	return buildData(customers)
}

func buildData(customers []model.Customer) ([]value.Phone, error) {
	var data []value.Phone

	for _, c := range customers {
		data = append(data, classify(c))
	}
	return data, nil
}

var countryMap = map[string]phoneItem{
	"237": {"Cameroon", regexp.MustCompile(`\(237\) ?[2368]\d{7,8}$`)},
	"251": {"Ethiopia", regexp.MustCompile(`\(251\) ?[1-59]\d{8}$`)},
	"212": {"Morocco", regexp.MustCompile(`\(212\) ?[5-9]\d{8}$`)},
	"258": {"Mozambique", regexp.MustCompile(`\(258\) ?[28]\d{7,8}$`)},
	"256": {"Uganda", regexp.MustCompile(`\(256\) ?\d{9}$`)},
}

func classify(customer model.Customer) value.Phone {
	prefix := strings.TrimFunc(customer.Phone[0:5], func(r rune) bool {
		return r == '(' || r == ')'
	})

	country := countryMap[prefix].name
	isValid := countryMap[prefix].isValid(customer.Phone)
	phone := strings.Fields(customer.Phone)[1]

	return value.Phone{
		Country:     country,
		CountryCode: "+" + prefix,
		State:       isValid,
		Phone:       phone,
	}
}

type phoneItem struct {
	name  string
	regex *regexp.Regexp
}

func (p phoneItem) isValid(number string) string {
	if p.regex.Match([]byte(number)) {
		return "OK"
	}
	return "NOK"
}

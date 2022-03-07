package service

import (
	"context"
	"github.com/rodrigodev/jumia-go/src/internal/phone/model"
	"github.com/rodrigodev/jumia-go/src/internal/phone/repository"
	"github.com/rodrigodev/jumia-go/src/internal/phone/value"
	"regexp"
	"strings"
)

// Container describes the service.
type Container interface {
	GetPhones(ctx context.Context) ([]model.Customer, error)
}

// PhoneService @impl Service static check
type PhoneService struct {
	repository repository.CustomerManager
}

func NewPhoneService(repository repository.CustomerManager) *PhoneService {
	return &PhoneService{
		repository,
	}
}

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

var CountryMap = map[string]phoneItem{
	"237": {"Cameroon", regexp.MustCompile("\\(237\\) ?[2368]\\d{7,8}$")},
	"251": {"Ethiopia", regexp.MustCompile("\\(251\\) ?[1-59]\\d{8}$")},
	"212": {"Morocco", regexp.MustCompile("\\(212\\) ?[5-9]\\d{8}$")},
	"258": {"Mozambique", regexp.MustCompile(" \\(258\\) ?[28]\\d{7,8}$")},
	"256": {"Uganda", regexp.MustCompile("\\(256\\) ?\\d{9}$")},
}

func classify(customer model.Customer) value.Phone {
	prefix := strings.TrimFunc(customer.Phone[0:5], func(r rune) bool {
		return r == '(' || r == ')'
	})

	country := CountryMap[prefix].name
	isValid := CountryMap[prefix].isValid(customer.Phone)
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

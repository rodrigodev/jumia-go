package transport

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rodrigodev/jumia-go/src/internal/phone/service"
	"github.com/rodrigodev/jumia-go/src/internal/phone/value"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const getPhoneURL = "/phone"

func TestHandler(t *testing.T) {
	t.Run("should implement http.Handler", func(t *testing.T) {
		var i interface{} = new(Handler)
		if _, ok := i.(http.Handler); !ok {
			t.Fail()
		}
	})

	t.Run("should return apply error", func(t *testing.T) {
		expectedErr := errors.New("some error")
		_, err := New(
			func(*Handler) error { return expectedErr },
		)
		require.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestGetPhonesHandler(t *testing.T) {
	t.Run("should successfully make request empty json - 200", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		phoneService := service.NewMockPhoneServiceContainer(ctrl)

		h, err := New(Phone(phoneService))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, getPhoneURL, nil)

		phoneService.
			EXPECT().
			GetPhones(gomock.Any()).
			Return([]value.Phone{}, nil).
			Times(1)

		h.ServeHTTP(w, r)
		res := w.Result()
		defer res.Body.Close()

		responseJson, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		require.JSONEq(t, "[]", string(responseJson))

		require.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("should successfully make request with data - 200", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		phoneService := service.NewMockPhoneServiceContainer(ctrl)

		h, err := New(Phone(phoneService))
		require.NoError(t, err)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, getPhoneURL, nil)

		phoneService.
			EXPECT().
			GetPhones(gomock.Any()).
			Return([]value.Phone{
				{
					Country:     "Morocco",
					State:       "OK",
					CountryCode: "+212",
					Phone:       "698054317",
				},
			}, nil).
			Times(1)

		h.ServeHTTP(w, r)
		res := w.Result()
		defer res.Body.Close()

		responseJson, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		require.JSONEq(t, `[{"country": "Morocco", "state": "OK", "country_code": "+212", "phone": "698054317"}]`, string(responseJson))

		require.Equal(t, http.StatusOK, res.StatusCode)
	})
}

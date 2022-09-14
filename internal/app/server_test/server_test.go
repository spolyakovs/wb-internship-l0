package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
)

func TestServerFunctions(t *testing.T) {
	if err := pub.PublishInvalid(); err != nil {
		t.Errorf("error while publishing invalid data: %v", err)
	}
	orderUIDRandom, err := pub.PublishRandomValid()
	if err != nil {
		t.Errorf("error while publishing random valid data: %v", err)
	}

	t.Run("Get without ID", func(t *testing.T) {
		res, err := serverGet("/")
		if err != nil {
			t.Errorf("unable to complete GET request without ID:%v", err)
			return
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("request without ID returned wrong status code:\n\tgot: %v\n\twanted: %v", res.StatusCode, http.StatusBadRequest)
			return
		}
	})

	t.Run("Get with empty ID", func(t *testing.T) {
		res, err := serverGet("/?id=")
		if err != nil {
			t.Errorf("unable to complete GET request with empty ID:%v", err)
			return
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("request with empty ID returned wrong status code:\n\tgot: %v\n\twanted: %v", res.StatusCode, http.StatusBadRequest)
			return
		}
	})

	t.Run("Get with non-existent ID", func(t *testing.T) {
		res, err := serverGet("/?id=not_exists")
		if err != nil {
			t.Errorf("unable to complete GET request with non-existent ID:%v", err)
			return
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("request with non-existent ID returned wrong status code:\n\tgot: %v\n\twanted: %v", res.StatusCode, http.StatusBadRequest)
			return
		}
	})

	t.Run("Get with valid ID", func(t *testing.T) {
		time.Sleep(200 * time.Millisecond)
		res, err := serverGet("/?id=" + orderUIDRandom)
		if err != nil {
			t.Errorf("unable to complete GET request with valid ID:%v", err)
			return
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("request with valid ID (%v) returned wrong status code:\n\tgot: %v\n\twanted: %v", orderUIDRandom, res.StatusCode, http.StatusOK)
			return
		}

		order := model.Order{}
		if err := json.NewDecoder(res.Body).Decode(&order); err != nil {
			t.Errorf("unable decode response body into model.Order:%v", err)
			return
		}

		if order.OrderUID != orderUIDRandom {
			t.Errorf("found order with another order\n\tgot: %+v\n\twanted: %+v", order.OrderUID, orderUIDRandom)
		}
	})

}

func serverGet(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	return w.Result(), nil
}

package server_test

import (
	"encoding/json"
	"net/http"
	"testing"

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
		res, err := http.Get(baseURL)
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
		res, err := http.Get(baseURL + "?id=")
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
		res, err := http.Get(baseURL + "?id=not_exists")
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
		res, err := http.Get(baseURL + "?id=" + orderUIDRandom)
		if err != nil {
			t.Errorf("unable to complete GET request with valid ID:%v", err)
			return
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("request with valid ID returned wrong status code:\n\tgot: %v\n\twanted: %v", res.StatusCode, http.StatusBadRequest)
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

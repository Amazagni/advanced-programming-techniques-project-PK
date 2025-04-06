package handlers

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestItems(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items", Items)

	req, _ := http.NewRequest("GET", "/items?limit=0&offset=0", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var actual []models.Item
	_ = json.Unmarshal(recorder.Body.Bytes(), &actual)

	expected := []models.Item{
		{
			Id:          1,
			Name:        "1",
			Description: "1",
			Quantity:    1,
			ImageURL:    "1.1",
		},
		{
			Id:          2,
			Name:        "2",
			Description: "2",
			Quantity:    2,
			ImageURL:    "2.2",
		}}

	assert.ElementsMatch(t, expected, actual)
}

func TestItemsLimitOffset(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items", Items)

	req, _ := http.NewRequest("GET", "/items?limit=1&offset=1", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var actual []models.Item
	_ = json.Unmarshal(recorder.Body.Bytes(), &actual)

	expected := []models.Item{
		{
			Id:          2,
			Name:        "2",
			Description: "2",
			Quantity:    2,
			ImageURL:    "2.2",
		}}

	assert.ElementsMatch(t, expected, actual)
}

func TestItemByID(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items/{id}", ItemByID).Methods("GET")

	req, _ := http.NewRequest("GET", "/items/2", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var actual models.Item
	_ = json.Unmarshal(recorder.Body.Bytes(), &actual)

	expected := models.Item{
		Id:          2,
		Name:        "2",
		Description: "2",
		Quantity:    2,
		ImageURL:    "2.2",
	}

	assert.Equal(t, expected, actual)
}

func TestItemByIDNotFound(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items/{id}", ItemByID).Methods("GET")

	req, _ := http.NewRequest("GET", "/items/999", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	var actual models.ErrorResponse
	_ = json.Unmarshal(recorder.Body.Bytes(), &actual)

	expected := models.ErrorResponse{
		Error:  "Item not found",
		Status: 404,
	}

	assert.Equal(t, expected, actual)
}

func TestCreateItem(t *testing.T) {
	newItem := models.Item{
		Id:          3,
		Name:        "Test Item",
		Description: "Test Description",
		Quantity:    10,
		ImageURL:    "test.jpg",
	}

	router := mux.NewRouter()
	router.HandleFunc("/items/create", CreateItem)

	data, _ := json.Marshal(newItem)
	req, _ := http.NewRequest("POST", "/items/create", bytes.NewBuffer(data))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestUpdateItem(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items/update/{id}", UpdateItem).Methods("POST")

	itemToUpdate := models.Item{
		Id:          5,
		Name:        "Updated Name",
		Description: "Updated Description",
		Quantity:    5,
		ImageURL:    "updated.url",
	}
	body, _ := json.Marshal(itemToUpdate)

	req, _ := http.NewRequest("POST", "/items/update/2", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, Data[1], itemToUpdate)
}

func TestUpdateNonExistingItem(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items/update/{id}", UpdateItem).Methods("POST")

	itemToUpdate := models.Item{
		Id:          5,
		Name:        "Updated Name",
		Description: "Updated Description",
		Quantity:    5,
		ImageURL:    "updated.url",
	}
	body, _ := json.Marshal(itemToUpdate)

	req, _ := http.NewRequest("POST", "/items/update/99", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 404, recorder.Code)
}

func TestDeleteItem(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items/delete/{id}", DeleteItem)

	req, _ := http.NewRequest("DELETE", "/items/delete/1", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 204, recorder.Code)
}

func TestDeleteNonExistentItem(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/items/delete/{id}", DeleteItem)

	req, _ := http.NewRequest("DELETE", "/items/delete/999", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 404, recorder.Code)
}

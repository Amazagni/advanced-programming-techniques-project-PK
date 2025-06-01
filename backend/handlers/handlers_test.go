package handlers

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestPaging(t *testing.T) {
	data := []models.Item{
		{Id: 1, Name: "Item 1"},
		{Id: 2, Name: "Item 2"},
		{Id: 3, Name: "Item 3"},
		{Id: 4, Name: "Item 4"},
	}

	tests := []struct {
		name     string
		limit    int
		offset   int
		expected []models.Item
	}{
		{
			name:     "Normal pagination",
			limit:    2,
			offset:   1,
			expected: data[1:3],
		},
		{
			name:     "Offset 0, limit 2",
			limit:    2,
			offset:   0,
			expected: data[0:2],
		},
		{
			name:     "Offset near end, limit exceeds",
			limit:    3,
			offset:   2,
			expected: data[2:4],
		},
		{
			name:     "Offset equals length",
			limit:    2,
			offset:   4,
			expected: []models.Item{},
		},
		{
			name:     "Offset > length",
			limit:    1,
			offset:   5,
			expected: []models.Item{},
		},
		{
			name:     "Negative limit",
			limit:    -1,
			offset:   1,
			expected: []models.Item{},
		},
		{
			name:     "Negative offset",
			limit:    1,
			offset:   -1,
			expected: []models.Item{},
		},
		{
			name:     "Limit 0 returns all",
			limit:    0,
			offset:   0,
			expected: data,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := paging(data, tt.limit, tt.offset)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestItems_Success(t *testing.T) {
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
			Name:        "Laptop",
			Description: "Can be used as a radiator too.",
			Quantity:    10,
			ImageURL:    "assets/images/placeholder.jpg",
		},
		{
			Id:          2,
			Name:        "Fridge",
			Description: "Idk how it got here",
			Quantity:    1,
			ImageURL:    "assets/images/placeholder.jpg",
		},
	}

	assert.ElementsMatch(t, expected, actual)
}

func setupCreateItem() {
	Data = []models.Item{}
}

func createMultipartRequest(t *testing.T, fields map[string]string) *http.Request {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for key, val := range fields {
		_ = writer.WriteField(key, val)
	}

	writer.Close()
	req := httptest.NewRequest("POST", "/items/create", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req
}

func TestCreateItem_SuccessWithoutImage(t *testing.T) {
	setupCreateItem()

	req := createMultipartRequest(t, map[string]string{
		"name":        "TestItem",
		"description": "TestDescription",
		"quantity":    "10",
	})

	rr := httptest.NewRecorder()
	CreateItem(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var created models.Item
	err := json.Unmarshal(rr.Body.Bytes(), &created)
	assert.NoError(t, err)
	assert.Equal(t, "TestItem", created.Name)
	assert.Equal(t, "TestDescription", created.Description)
	assert.Equal(t, int32(10), created.Quantity)
	assert.Equal(t, "", created.ImageURL) // Brak zdjęcia
}

func TestCreateItem_OPTIONSMethod(t *testing.T) {
	setupCreateItem()

	req := httptest.NewRequest("OPTIONS", "/items/create", nil)
	rr := httptest.NewRecorder()

	CreateItem(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func setupDeleteItemData() {
	Data = []models.Item{
		{
			Id:          1,
			Name:        "Test Item",
			Description: "Test Desc",
			Quantity:    5,
			ImageURL:    "assets/images/test.jpg",
		},
		{
			Id:          2,
			Name:        "Another Item",
			Description: "Another Desc",
			Quantity:    3,
			ImageURL:    "assets/images/test2.jpg",
		},
	}
}

func setupItemByID() {
	Data = []models.Item{
		{
			Id:          1,
			Name:        "Item1",
			Description: "Test item 1",
			Quantity:    5,
			ImageURL:    "",
		},
		{
			Id:          2,
			Name:        "Item2",
			Description: "Test item 2",
			Quantity:    10,
			ImageURL:    "",
		},
	}
}

func TestItemByID_Found(t *testing.T) {
	setupItemByID()

	req := httptest.NewRequest("GET", "/items/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()

	ItemByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var item models.Item
	err := json.Unmarshal(rr.Body.Bytes(), &item)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), item.Id)
	assert.Equal(t, "Item1", item.Name)
}

func TestItemByID_NotFound(t *testing.T) {
	setupItemByID()

	req := httptest.NewRequest("GET", "/items/999", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "999"})
	rr := httptest.NewRecorder()

	ItemByID(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "Item not found")
}

func TestItemByID_OPTIONS(t *testing.T) {
	req := httptest.NewRequest("OPTIONS", "/items/1", nil)
	rr := httptest.NewRecorder()

	ItemByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteItem_Success(t *testing.T) {
	setupDeleteItemData()

	req, err := http.NewRequest("DELETE", "/items/delete/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items/delete/{id}", DeleteItem).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Len(t, Data, 1)
	assert.NotEqual(t, int32(1), Data[0].Id)
}

func TestDeleteItem_NotFound(t *testing.T) {
	setupDeleteItemData()

	req, err := http.NewRequest("DELETE", "/items/delete/999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items/delete/{id}", DeleteItem).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "Can't delete item")
	assert.Len(t, Data, 2)
}

func TestDeleteItem_OptionsMethod(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "/items/delete/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/items/delete/{id}", DeleteItem).Methods("DELETE", "OPTIONS")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

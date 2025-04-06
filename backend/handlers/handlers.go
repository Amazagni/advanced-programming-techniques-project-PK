package handlers

import (
	"backend/logger"
	"backend/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"slices"
	"strconv"
)

var (
	Data = []models.Item{
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
)

func WriteError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)

	errorResponse := models.ErrorResponse{
		Error:  message,
		Status: statusCode,
	}

	data, _ := json.Marshal(errorResponse)
	_, _ = w.Write(data)
}

func paging(data []models.Item, limit, offset int) []models.Item {
	n := len(data)
	if limit == 0 {
		return data
	}
	if limit < 0 || offset < 0 || offset >= n {
		return make([]models.Item, 0)
	}

	end := offset + limit
	if end > n {
		end = n
	}
	return data[offset:end]
}

func Items(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitStr := query.Get("limit")
	offsetStr := query.Get("offset")
	var limit int
	var offset int
	var err error

	if limitStr == "" {
		WriteError(w, "Bad limit", 404)
		return
	}
	limit, err = strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		logger.Error("Bad limit")
		WriteError(w, "Can't parse limit", 501)
		return
	}

	if offsetStr == "" {
		WriteError(w, "Bad offset", 404)
		return
	}
	offset, err = strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		logger.Error("Bad offset")
		WriteError(w, "Can't parse offset", 501)
		return
	}
	dataLimited := paging(Data, limit, offset)

	dataByte, _ := json.Marshal(dataLimited)
	_, _ = w.Write(dataByte)
}

func ItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 16)
	id := int32(i)

	for _, item := range Data {
		if item.Id == id {
			dataByte, _ := json.Marshal(item)
			_, _ = w.Write(dataByte)
			w.WriteHeader(200)
			return
		}
	}
	WriteError(w, "Item not found", 404)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	body, err := io.ReadAll(r.Body)
	err = json.Unmarshal(body, &item)
	if err != nil {
		logger.Error("Can't create item")
		WriteError(w, "Can't create Item", 404)
	}
	Data = append(Data, item)
	w.WriteHeader(200)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 16)
	id := int32(i)

	requestItem := models.Item{}
	body, err := io.ReadAll(r.Body)
	err = json.Unmarshal(body, &requestItem)
	if err != nil {
		logger.Error("Can't update item")
		WriteError(w, "Can't update item", 404)
	}

	for idx, item := range Data {
		if item.Id == id {
			Data[idx] = requestItem
			w.WriteHeader(200)
			return
		}
	}
	WriteError(w, "Can't update item", 404)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 16)
	id := int32(i)
	for idx, item := range Data {
		if item.Id == id {
			Data = slices.Concat(Data[:idx], Data[idx+1:])
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	WriteError(w, "Can't delete item", 404)
}

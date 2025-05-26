package handlers

import (
	_ "backend/docs"
	"backend/logger"
	"backend/models"
	"encoding/json"
	"io"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	Data = []models.Item{
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

// Items godoc
// @Summary Get all items
// @Description Returns all available items
// @Tags items
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Item
// @Router /items [get]
func Items(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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

// ItemByID godoc
// @Summary Get item by ID
// @Description Returns a single item by ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Failure 404 {object} models.ErrorResponse
// @Router /items/{id} [get]
func ItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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

// CreateItem godoc
// @Summary Create a new item
// @Description Creates a new item with form data
// @Tags items
// @Accept multipart/form-data
// @Produce  json
// @Param name formData string true "Item name"
// @Param description formData string true "Item description"
// @Param quantity formData int true "Item quantity"
// @Param image formData file false "Item image"
// @Success 201 {object} models.Item
// @Failure 400 {object} models.ErrorResponse
// @Router /items/create [post]
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for the form
	if err != nil {
		logger.Error("Can't parse form")
		WriteError(w, "Can't parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	quantityStr := r.FormValue("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		logger.Error("Invalid quantity")
		WriteError(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		logger.Error("Error retrieving image file")
		WriteError(w, "Error retrieving image file", http.StatusBadRequest)
		return
	}
	var imageURL string
	if file != nil {
		defer file.Close()
		// Save the image (you'll need to implement your own logic here)
		imageURL = header.Filename // Or generate a unique filename
	}

	newItem := models.Item{
		Name:        name,
		Description: description,
		Quantity:    int32(quantity),
		ImageURL:    imageURL,
	}

	Data = append(Data, newItem)
	w.WriteHeader(http.StatusCreated)
	dataByte, _ := json.Marshal(newItem)
	_, _ = w.Write(dataByte)

	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	logger.Error("Error reading request body")
	// 	WriteError(w, "Error reading request body", http.StatusBadRequest)
	// 	return
	// }

	// item := models.Item{}
	// err = json.Unmarshal(body, &item)
	// logger.Info("Received request body: " + string(body)) // Log the raw body
	// if err != nil {
	// 	logger.Error("Can't create item")
	// 	WriteError(w, "Can't create Item. Error: "+err.Error(), 400)
	// }

	// Data = append(Data, item)
	// w.WriteHeader(200)
}

// UpdateItem godoc
// @Summary Update item quantity
// @Description Updates quantity of a specific item
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param quantity body int true "New quantity"
// @Success 200 {object} models.Item
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /items/update/{id} [put]
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	vars := mux.Vars(r)
	i, err := strconv.ParseInt(vars["id"], 10, 16)
	if err != nil {
		logger.Error("Invalid item ID")
		WriteError(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	id := int32(i)
	logger.InfOf("Updating quantity for item ID: %d", id)

	var requestBody map[string]interface{} // Expecting a JSON with { "quantity": newQuantity }
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Can't read request body")
		WriteError(w, "Can't read request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		logger.Error("Can't parse request body")
		WriteError(w, "Can't parse request body", http.StatusBadRequest)
		return
	}

	newQuantityFloat, ok := requestBody["Quantity"].(float64)
	if !ok {
		logger.Error("Invalid quantity format")
		WriteError(w, "Invalid quantity format", http.StatusBadRequest)
		return
	}
	newQuantity := int32(newQuantityFloat)
	logger.InfOf("New quantity received: %d", newQuantity)

	for idx, item := range Data {
		logger.InfOf("Checking item with ID: %d", item.Id)
		if item.Id == id {
			logger.InfOf("Found item with ID: %d, updating quantity to: %d", item.Id, newQuantity)
			Data[idx].Quantity = newQuantity
			logger.InfOf("Updated quantity in Data slice: %d", Data[idx].Quantity)

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(Data[idx])
			if err != nil {
				logger.Error("Error encoding updated item")
			}
			logger.InfOf("Encoded and sent updated item: %+v", Data[idx])
			return
		}
	}

	// requestItem := models.Item{}
	// body, err := io.ReadAll(r.Body)
	// err = json.Unmarshal(body, &requestItem)
	// if err != nil {
	// 	logger.Error("Can't update item")
	// 	WriteError(w, "Can't update item", 404)
	// }

	// for idx, item := range Data {
	// 	if item.Id == id {
	// 		Data[idx] = requestItem
	// 		w.WriteHeader(200)
	// 		return
	// 	}
	// }
	WriteError(w, "Can't update item", 404)
}

// DeleteItem godoc
// @Summary Delete an item
// @Description Deletes item by ID
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Router /items/delete/{id} [delete]
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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

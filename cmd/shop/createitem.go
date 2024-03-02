package main

import (
    "ecommerce/pkg/shop/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func (h handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating an item...")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	var item model.Item
	json.Unmarshal(body, &item)

	if result := h.DB.Create(&item); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

	fmt.Println("item has been created")
}

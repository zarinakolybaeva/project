package main

import (
	"ecommerce/pkg/shop/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h handler) DeleteItemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	ErrorHandler(err)

	var item model.Item

	if result := h.DB.Where("id = ?", id).Delete(&item); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("item is deleted")
}

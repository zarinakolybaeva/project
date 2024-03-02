package main

import (
	"ecommerce/pkg/shop/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (h handler) UpdateItemById(w http.ResponseWriter, r *http.Request) {


	var item model.Item

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	ErrorHandler(err)

	json.Unmarshal(body, &item)

	fmt.Println(item)

	if result := h.DB.Model(&item).Updates(&item); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(item)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)

	fmt.Println("item is updated")
}

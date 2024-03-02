package main
import (
	"encoding/json"
	"ecommerce/pkg/shop/model"
	"fmt"
	"net/http"
)

func (h handler ) GetAllItems(w http.ResponseWriter, r *http.Request) {
	var items []model.Item
	if result := h.DB.Find(&items); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}


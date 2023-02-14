package controller

import (
	"awesomeProject/kofeslivka/pkg/entity"
	"awesomeProject/kofeslivka/pkg/usecase"
	"encoding/json"
	"net/http"
)

type Controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := &entity.RawMaterial{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
		}
		id, err := c.usecase.CreateUser(user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
		}
		result := map[string]int64{"id": id}
		response, err := json.Marshal(result)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}
}

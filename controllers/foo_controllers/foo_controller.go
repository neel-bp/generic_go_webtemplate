package foo_controllers

import (
	"encoding/json"
	"net/http"

	"wagie.com/wageslavery/entity/foo_entity"
	fooservices "wagie.com/wageslavery/serviices/foo_services"
	"wagie.com/wageslavery/viewmodels"
)

type FooController struct {
	FooSvc foo_entity.IFooService
}

func New() *FooController {
	return &FooController{
		FooSvc: fooservices.New(),
	}
}

func (fc *FooController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res, err := fc.FooSvc.Get()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

func (fc *FooController) Create(w http.ResponseWriter, r *http.Request) {
	var body viewmodels.FooGet
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	err = fc.FooSvc.Create(body)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "success"})
}

package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-employee/models"
	"go-rest-employee/pkg/api/utils"
	"go-rest-employee/pkg/service"
	"net/http"
	"strconv"
)

func Create(ctx utils.MyContext, service service.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee models.Employee
		if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := service.Create(employee)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"id": id,
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetByCompany(ctx utils.MyContext, service service.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companyId, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		employees, err := service.GetByCompany(companyId)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(employees); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetByDepartment(ctx utils.MyContext, service service.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		departmentName := mux.Vars(r)["departmentName"]

		employees, err := service.GetByDepartment(departmentName)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(employees); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Update(ctx utils.MyContext, service service.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		var input models.Employee
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = service.Update(id, input); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Delete(ctx utils.MyContext, service service.EmployeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = service.Delete(id); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(utils.StatusResponse{Status: "ok"}); err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

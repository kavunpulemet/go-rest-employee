package api

import (
	"context"
	"github.com/gorilla/mux"
	"go-rest-employee/pkg/api/handler"
	"go-rest-employee/pkg/api/utils"
	"go-rest-employee/pkg/service"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter()
	return &Server{
		httpServer: &http.Server{
			Addr:           ":80",
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			Handler:        router,
		},
		router: router,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) HandleEmployees(service service.EmployeeService, ctx utils.MyContext) {
	s.router.HandleFunc("/employees/", handler.Create(service, ctx)).Methods(http.MethodPost)
	s.router.HandleFunc("/employees/companies/{companyId}/", handler.GetByCompany(service, ctx)).Methods(http.MethodGet)
	s.router.HandleFunc("/employees/departments/{departmentName}/", handler.GetByDepartment(service, ctx)).Methods(http.MethodGet)
	s.router.HandleFunc("/employees/{id}/", handler.Update(service, ctx)).Methods(http.MethodPut)
	s.router.HandleFunc("/employees/{id}/", handler.Delete(service, ctx)).Methods(http.MethodDelete)
}

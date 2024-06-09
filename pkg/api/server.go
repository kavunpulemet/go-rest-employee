package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go-rest-employee/pkg/api/handler"
	"go-rest-employee/pkg/api/middlewares"
	"go-rest-employee/pkg/api/utils"
	"go-rest-employee/pkg/service"
	"net/http"
	"time"
)

const (
	maxHeaderBytes = 1 << 20 // 1 MB
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(ctx utils.MyContext) *Server {
	router := mux.NewRouter()

	wrappedRouter := middlewares.RecoveryMiddleware(ctx, router)

	return &Server{
		httpServer: &http.Server{
			Addr:           viper.GetString("db"),
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			Handler:        wrappedRouter,
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

func (s *Server) HandleEmployees(ctx utils.MyContext, service service.EmployeeService) {
	s.router.HandleFunc("/employees/", handler.Create(ctx, service)).Methods(http.MethodPost)
	s.router.HandleFunc("/employees/companies/{companyId}/", handler.GetByCompany(ctx, service)).Methods(http.MethodGet)
	s.router.HandleFunc("/employees/departments/{departmentName}/", handler.GetByDepartment(ctx, service)).Methods(http.MethodGet)
	s.router.HandleFunc("/employees/{id}/", handler.Update(ctx, service)).Methods(http.MethodPut)
	s.router.HandleFunc("/employees/{id}/", handler.Delete(ctx, service)).Methods(http.MethodDelete)
}

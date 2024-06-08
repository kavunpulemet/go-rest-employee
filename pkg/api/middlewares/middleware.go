package middlewares

import (
	"go-rest-employee/pkg/api/utils"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware(ctx utils.MyContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic occurred: %v\n%s", err, debug.Stack())
				utils.NewErrorResponse(ctx, w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

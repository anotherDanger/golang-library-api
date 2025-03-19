package middleware

import (
	"encoding/json"
	"golang-library-api/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(Handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: Handler,
	}
}

func (auth *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apikey := r.Header.Get("X-API-Key")
	if apikey == "SECRET" {

		auth.Handler.ServeHTTP(w, r)
	} else {
		response := &web.WebResponse{
			Code:   401,
			Status: "unauthorized",
			Data:   nil,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(response)

	}
}

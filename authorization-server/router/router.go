package router

import (
	"authorization-server/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupDefault(r chi.Router) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func SetupOauth(r chi.Router, ho handler.Oauth) {
	r.Get("/authorize", ho.Authorize)
	r.Get("/certs/jwks", ho.CertificateJWKS)
}

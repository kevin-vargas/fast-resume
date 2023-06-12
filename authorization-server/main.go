package main

import (
	"authorization-server/certs"
	"authorization-server/cmd/config"
	"authorization-server/handler"
	"authorization-server/router"
	"authorization-server/service"
	"authorization-server/service/jwt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

const (
	cert_file = "cert/cert.pem"
	key_file  = "cert/key.pem"
)

func main() {
	cfg := config.Make()

	pk, err := certs.MakePrivateKey(cfg.PrivateKey)
	if err != nil {
		os.Exit(1)
	}
	var oauthConfig = oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL:  cfg.TokenURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	tm := jwt.NewTokenManager(pk, cfg.Kid, cfg.TokenDuration)
	so := service.NewOauth(oauthConfig, tm)
	ho := handler.NewOauth(so)
	r := chi.NewRouter()
	router.SetupDefault(r)
	router.SetupOauth(r, ho)
	if cfg.Secure == "true" {
		err = http.ListenAndServeTLS(cfg.Port, cert_file, key_file, r)
	} else {
		err = http.ListenAndServe(cfg.Port, r)
	}
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

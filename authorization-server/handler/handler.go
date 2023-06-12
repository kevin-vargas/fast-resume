package handler

import (
	"authorization-server/service"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Oauth interface {
	CertificateJWKS(http.ResponseWriter, *http.Request)
	Authorize(http.ResponseWriter, *http.Request)
}

type oauth struct {
	svc service.Oauth
}

func (o *oauth) CertificateJWKS(w http.ResponseWriter, r *http.Request) {
	cert, err := o.svc.GetCertificateJWKS()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	raw, err := json.Marshal(cert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(raw)
}

func (o *oauth) Authorize(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	slackToken, err := o.svc.Exchange(code)
	if err != nil {
		log.Printf("On exchange token: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := o.svc.GenerateJWT(slackToken)
	if err != nil {
		log.Printf("On generate jwt token: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	origin, err := url.QueryUnescape(r.URL.Query().Get("state"))
	if err != nil {
		log.Printf("On query unescape: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "logged",
		Value: "true",
	})
	http.Redirect(w, r, origin, http.StatusSeeOther)
}

func NewOauth(svc service.Oauth) Oauth {
	return &oauth{
		svc: svc,
	}
}

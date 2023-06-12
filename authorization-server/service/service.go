package service

import (
	"authorization-server/entity"
	"authorization-server/service/jwt"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	Alg = "RS256"
	Kty = "RSA"
	Use = "sig"
	E   = "AQAB"
)

type Oauth interface {
	GetCertificateJWKS() (*entity.Jwks, error)
	GenerateJWT(*entity.SlackToken) (string, error)
	Exchange(string) (*entity.SlackToken, error)
}

type oauth struct {
	tm jwt.TokenManager
	oauth2.Config
}

func (o *oauth) Exchange(code string) (*entity.SlackToken, error) {
	v := url.Values{
		"client_id":     {o.ClientID},
		"client_secret": {o.ClientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}
	res, err := http.PostForm(o.Endpoint.TokenURL, v)
	if err != nil {
		return nil, err
	}
	var e entity.SlackToken
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		return nil, err
	}
	if !e.Ok {
		return nil, fmt.Errorf("invalid token %+v, %s", e, e.Err)
	}
	return &e, nil
}

func (o *oauth) GenerateJWT(sk *entity.SlackToken) (string, error) {
	c := entity.Claim{
		Token: sk.AuthedUser.AccessToken,
	}
	return o.tm.Generate(c)
}

func (o *oauth) GetCertificateJWKS() (*entity.Jwks, error) {
	pk := o.tm.PrivateKey().PublicKey
	n := base64.StdEncoding.EncodeToString(pk.N.Bytes())
	return &entity.Jwks{
		Keys: []entity.JwksKey{
			{
				Kid: o.tm.Kid(),
				Alg: Alg,
				Kty: Kty,
				Use: Use,
				N:   n,
				E:   E,
			},
		},
	}, nil
}

func NewOauth(c oauth2.Config, tm jwt.TokenManager) Oauth {
	return &oauth{
		tm:     tm,
		Config: c,
	}
}

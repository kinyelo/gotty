package server

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/pquerna/otp/totp"
)

type Auth2Fa struct {
	creds map[string]Credentials
}

type Credentials struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	OtpSecret string `json:"otpSecret"`
}

func NewAuth2Fa() (*Auth2Fa, error) {
	data, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, errors.New("could not open credentials.json")
	}
	var list []Credentials
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, errors.New("can't parse json from credentials.json")
	}
	credsMap := make(map[string]Credentials)
	for _, item := range list {
		credsMap[item.Username] = item
	}
	return &Auth2Fa{
		creds: credsMap,
	}, nil
}

func (auth *Auth2Fa) Valid(payload []byte) (bool, error) {
	p := strings.Split(string(payload), ":")
	if len(p[1]) <= 6 {
		return false, errors.New("wrong credentials: too short")
	}
	username, pToken := p[0], p[1]
	password := pToken[0 : len(pToken)-6]
	otpToken := pToken[len(pToken)-6:]

	creds, ok := auth.creds[username]
	if !ok {
		return false, errors.New("wrong credentials: no user found")
	}
	if password != creds.Password {
		return false, errors.New("wrong credentials: password incorrect")
	}
	return totp.Validate(otpToken, creds.OtpSecret), nil
}

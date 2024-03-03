package server

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
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
	configName := "/tmp/credentials.json"
	data, err := os.ReadFile(configName)
	os.Remove(configName)
	if err != nil {
		return nil, errors.New("could not open " + configName)
	}
	var list []Credentials
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, errors.New("can't parse json from " + configName)
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

	if getHash(password) != creds.Password {
		return false, errors.New("wrong credentials: password incorrect")
	}
	return totp.Validate(otpToken, creds.OtpSecret), nil
}

func getHash(passw string) string {
	h := sha256.New()
	h.Write([]byte(passw))
	return fmt.Sprintf("%x", h.Sum(nil))
}

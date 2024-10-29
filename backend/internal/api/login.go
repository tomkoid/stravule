package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type User struct {
	cislo             string
	dochazka          string
	editaceJidelnicku int
	email             string
	id                string
	jmeno             string
	konto             string
	limit             string
	mena              string
	nazevJidelny      string
	pasivni           bool
	pocetJidel        bool
	podrobnosti       bool
	prihlaska         bool
	slabeHeslo        bool
	text              string
	verze             string
	vicenasobny       bool
	vs                string
	vydej             int
}

type LoginResponse struct {
	Betatest        bool   `json:"betatest"`
	Cislo           string `json:"cislo"`
	IgnoreCert      bool   `json:"ignoreCert"`
	Jmeno           string `json:"jmeno"`
	S5url           string `json:"s5url"`
	SID             string `json:"sid"`
	Uzivatel        User   `json:"uzivatel"`
	Verze           string `json:"verze"`
	ZustatPrihlasen bool   `json:"zustatPrihlasen"`
}

type LoginRequest struct {
	Cislo           string `json:"cislo"`
	Jmeno           string `json:"jmeno"`
	Heslo           string `json:"heslo"`
	ZustatPrihlasen bool   `json:"zustatPrihlasen"`
	Environment     string `json:"environment"`
	Lang            string `json:"lang"`
}

// Returns SID and the canteen number
func Login(username string, password string, canteen string) (string, string, error) {
	requestBody := LoginRequest{
		Cislo:           canteen,
		Jmeno:           username,
		Heslo:           password,
		ZustatPrihlasen: true,
		Environment:     "W",
		Lang:            "EN",
	}

	requestBodyJson, err := json.Marshal(requestBody)
	if err != nil {
		return "", "", err
	}

	r, err := http.NewRequest("POST", "https://app.strava.cz/api/login", bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return "", "", err
	}

	r.Header.Set("accept", "*/*")
	r.Header.Set("accept-language", "en-US,en;q=0.9")
	r.Header.Set("content-type", "text/plain;charset=UTF-8")
	r.Header.Set("priority", "u=1, i")
	r.Header.Set("origin", "https://app.strava.cz")
	r.Header.Set("referer", "https://app.strava.cz/en/prihlasit-se?jidelna")
	r.Header.Set("referrer", "https://app.strava.cz/en/prihlasit-se?jidelna")
	r.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")

	res, err := client.Do(r)
	if err != nil {
		return "", "", err
	}

	if res.StatusCode == 555 {
		return "", "", errors.New("Wrong credentials, try again with the right credentials")
	}

	if res.StatusCode != 200 {
		return "", "", errors.New(fmt.Sprintf("Status code: %d", res.StatusCode))
	}

	response := LoginResponse{}
	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		return "", "", derr
	}

	return response.SID, response.Cislo, nil
}

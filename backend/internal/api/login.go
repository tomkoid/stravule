package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"codeberg.org/tomkoid/stravule/backend/db"
	"codeberg.org/tomkoid/stravule/backend/internal/config"
	"codeberg.org/tomkoid/stravule/backend/internal/database"
	"codeberg.org/tomkoid/stravule/backend/internal/utils"
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

type LoginReturn struct {
	SID      string `json:"sid"`
	Canteen  string `json:"canteen"`
	UserHash string `json:"user_hash"`
}

// Returns SID and the canteen number
func Login(username string, password string, canteen string) (*LoginReturn, error) {
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
		return nil, err
	}

	r, err := http.NewRequest("POST", "https://app.strava.cz/api/login", bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	if res.StatusCode == 555 {
		return nil, errors.New("Wrong credentials, try again with the right credentials")
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d", res.StatusCode))
	}

	response := LoginResponse{}
	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		return nil, derr
	}

	ctx := context.Background()
	userHash := utils.GetUserHash(username, password, canteen)

	betatesterErrorMsg := "Tento uživatel není betatesterem. Zeptej se administrátora této instance a napiš mu tvůj kód, aby ti povolil přístup. Tvůj kód je: " + userHash

	// insert new user into db if it doesn't exist
	user, err := database.DB.GetUser(ctx, userHash)
	if err != nil {
		newUser, err := database.DB.CreateUser(ctx, db.CreateUserParams{
			UserHash: userHash,
			Sid:      response.SID,
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("db: created new user with hash", userHash)
		}

		if config.Cfg.BetatestersOnly && !newUser.IsBetaTester {
			return nil, errors.New(betatesterErrorMsg)
		}
	} else {
		if config.Cfg.BetatestersOnly && !user.IsBetaTester {
			return nil, errors.New(betatesterErrorMsg)
		}
	}

	data := LoginReturn{
		SID:      response.SID,
		Canteen:  response.Cislo,
		UserHash: userHash,
	}

	return &data, nil
}

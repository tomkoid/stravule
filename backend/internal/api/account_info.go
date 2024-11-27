package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type accountInfoRequest struct {
	CheckVersion     bool   `json:"checkVersion"`
	Cislo            string `json:"cislo"`
	FrontendFunction string `json:"frontendFunction"`
	GetText          bool   `json:"getText"`
	IgnoreCert       string `json:"ignoreCert"`
	Lang             string `json:"lang"`
	ResetTables      bool   `json:"resetTables"`
	SID              string `json:"sid"`
	URL              string `json:"url"`
}

type AccountInfoResponse struct {
	Konto             string `json:"konto"`
	VS                string `json:"vs"`
	Jmeno             string `json:"jmeno"`
	Email             string `json:"email"`
	Vydej             int    `json:"vydej"`
	Limit             string `json:"limit"`
	Vicenasobny       bool   `json:"vicenasobny"`
	PocetJidel        int    `json:"pocetJidel"`
	Podrobnosti       int    `json:"podrobnosti"`
	EditaceJidelnicku int    `json:"editaceJidelnicku"`
	ID                string `json:"id"`
	Cislo             string `json:"cislo"`
	BakalarID         string `json:"bakalarId"`
	Prihlaska         int    `json:"prihlaska"`
	Pasivni           bool   `json:"pasivni"`
	Verze             string `json:"verze"`
	Text              string `json:"text"`
	Mena              string `json:"mena"`
	Dochazka          string `json:"dochazka"`
	NazevJidelny      string `json:"nazevJidelny"`
}

func GetAccountInfo(sid string, canteen string) (*AccountInfoResponse, error) {
	requestBody := accountInfoRequest{
		CheckVersion:     true,
		Cislo:            canteen,
		FrontendFunction: "refreshInformations",
		GetText:          true,
		Lang:             "EN",
		ResetTables:      true,
		IgnoreCert:       "false", // lol
		SID:              sid,
		URL:              "",
	}

	requestBodyJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("POST", "https://app.strava.cz/api/nactiVlastnostiPA", bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return nil, err
	}

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	// read body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 555 {
		return nil, errors.New("Wrong credentials, try again with the right credentials")
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d", res.StatusCode))
	}

	var response AccountInfoResponse

	buf := bytes.NewBuffer(bodyBytes)
	derr := json.NewDecoder(buf).Decode(&response)
	if derr != nil {
		return nil, derr
	}

	return &response, nil
}

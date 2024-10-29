package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type orderRequest struct {
	Cislo      string `json:"cislo"`
	SID        string `json:"sid"`
	Lang       string `json:"lang"`
	S5url      string `json:"s5url"`
	Podminka   string `json:"podminka"`
	Konto      int    `json:"konto"`
	IgnoreCert string `json:"ignoreCert"`
}

type order struct {
	Datum           string     `json:"datum"`
	Druh            string     `json:"druh"`
	Chod            string     `json:"chod"`
	Pocet           int        `json:"pocet"`
	Cena            string     `json:"cena"`
	Omezeni         string     `json:"omezeni"`
	Polevka         string     `json:"polevka"`
	Alergjson       []string   `json:"alergjson"`
	ID              int        `json:"id"`
	Nazev           string     `json:"nazev"`
	DruhPopis       string     `json:"druh_popis"`
	DruhChod        string     `json:"druh_chod"`
	Databaze        string     `json:"databaze"`
	DelsiPopis      string     `json:"delsiPopis"`
	AlergenyText    string     `json:"alergeny_text"`
	Alergeny        [][]string `json:"alergeny"`
	ZkratkaProduktu string     `json:"zkratkaProduktu"`
	MultipleNazev   string     `json:"multipleNazev"`
	Version         int        `json:"version"`
	CasKonec        string     `json:"casKonec"`
	CasOdhlaseni    string     `json:"casOdhlaseni"`
}

// Returns a list of orders
func Orders(sid string, canteen string) ([][]order, error) {
	requestBody := orderRequest{
		Cislo:      canteen,
		Lang:       "EN",
		S5url:      "",
		Podminka:   "",
		Konto:      0,
		IgnoreCert: "false", // lol
		SID:        sid,
	}

	requestBodyJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("POST", "https://app.strava.cz/api/objednavky", bytes.NewBuffer(requestBodyJson))
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
		panic("xd")
	}
	bodyString := string(bodyBytes)
	// fmt.Println(bodyString)
	_ = bodyString

	if res.StatusCode == 555 {
		return nil, errors.New("Wrong credentials, try again with the right credentials")
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code: %d", res.StatusCode))
	}

	var response map[string][]order

	buf := bytes.NewBuffer(bodyBytes)
	derr := json.NewDecoder(buf).Decode(&response)
	if derr != nil {
		return nil, derr
	}
	//
	var orders [][]order
	//
	// // for each table in the json response
	// for _, table := range response {
	// 	orders = append(orders, table)
	// }

	for tableName, tableOrders := range response {
		// fmt.Println("==============xd================")
		// fmt.Println(tableName, orders)
		_ = tableName
		orders = append(orders, tableOrders)
	}

	return orders, nil
	// return response.SID, response.Cislo, nil
}

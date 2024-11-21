package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SendOrdersRequest struct {
	Cislo      string `json:"cislo"`
	SID        string `json:"sid"`
	Url        string `json:"url"`
	Xml        string `json:"xml"`
	Lang       string `json:"lang"`
	IgnoreCert string `json:"ignoreCert"`
}

func addOrder(datum string, druh string, pocet int) string {
	// Parse the input date using the appropriate layout
	parsedDate, err := time.Parse("02.01.2006", datum)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	// Format the date to the desired format
	outputDate := parsedDate.Format("2006-01-02")

	orderStr := fmt.Sprintf("&lt;rozpisobjednavek&gt;\n&lt;datum&gt;%s&lt;/datum&gt;\n&lt;druh&gt;%s&lt;/druh&gt;\n&lt;pocet&gt;%d&lt;/pocet&gt;\n&lt;/rozpisobjednavek&gt;", outputDate, druh, pocet)
	return orderStr
}

func SendOrders(request SendOrdersRequest, userHash string) error {
	// get orders
	pickedOrders, originalOrders, err := PickOrders(request.SID, request.Cislo, userHash)
	if err != nil {
		return err
	}

	ordersStr := ""

	// get changes
	for i := range pickedOrders {
		for j := range pickedOrders[i] {
			if pickedOrders[i][j].Pocet != originalOrders[i][j].Pocet {
				ordersStr += addOrder(pickedOrders[i][j].Datum, pickedOrders[i][j].Druh, pickedOrders[i][j].Pocet)
			}
		}
	}

	request.Xml = fmt.Sprintf("&lt;?xml version=\"1.0\" encoding=\"UTF-8\"?&gt;\n  &lt;VFPData&gt;\n     &lt;xsd:schema xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:msdata=\"urn:schemas-microsoft-com:xml-msdata\" id=\"VFPData\"&gt;\n        &lt;xsd:element name=\"VFPData\" msdata:IsDataSet=\"true\"&gt;\n           &lt;xsd:complexType&gt;\n              &lt;xsd:choice maxOccurs=\"unbounded\"&gt;\n                 &lt;xsd:element name=\"rozpisobjednavek\" minOccurs=\"0\" maxOccurs=\"unbounded\"&gt;\n                    &lt;xsd:complexType&gt;\n                       &lt;xsd:sequence&gt;\n                          &lt;xsd:element name=\"datum\" type=\"xsd:date\" /&gt;\n                          &lt;xsd:element name=\"druh\"&gt;\n                             &lt;xsd:simpleType&gt;\n                                &lt;xsd:restriction base=\"xsd:string\"&gt;\n                                   &lt;xsd:maxLength value=\"1\" /&gt;\n                                &lt;/xsd:restriction&gt;\n                             &lt;/xsd:simpleType&gt;\n                          &lt;/xsd:element&gt;\n                          &lt;xsd:element name=\"pocet\"&gt;\n                             &lt;xsd:simpleType&gt;\n                                &lt;xsd:restriction base=\"xsd:decimal\"&gt;\n                                   &lt;xsd:totalDigits value=\"5\" /&gt;\n                                   &lt;xsd:fractionDigits value=\"0\" /&gt;\n                                &lt;/xsd:restriction&gt;\n                             &lt;/xsd:simpleType&gt;\n                          &lt;/xsd:element&gt;\n                       &lt;/xsd:sequence&gt;\n                    &lt;/xsd:complexType&gt;\n                 &lt;/xsd:element&gt;\n              &lt;/xsd:choice&gt;\n              &lt;xsd:anyAttribute namespace=\"http://www.w3.org/XML/1998/namespace\" processContents=\"lax\" /&gt;\n           &lt;/xsd:complexType&gt;\n        &lt;/xsd:element&gt;\n     &lt;/xsd:schema&gt;%s\n&lt;/VFPData&gt;\n", ordersStr) // seriously, why do i have to do this

	requestJson, err := json.Marshal(request)
	req, err := http.NewRequest("POST", "https://app.strava.cz/api/saveOrders", bytes.NewReader(requestJson))
	req.Header.Set("Content-Type", "text/plain") // whyyyyyyyyy

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("send orders: status code: %d", resp.StatusCode)
	}

	return nil
}

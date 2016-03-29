package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//in ESRI REST, you want the features[] - this is where results are stored (other info is metadata but starting to add it).
type query struct {
	Features         []features
	GeometryType     string
	DisplayFieldName string
}

//Features has another [] of attributes. You can either hardcode all the values like "OBJECTID string" or if you want more generic, you can map it. then below you can iterate.

type features struct {
	Attributes map[string]interface{}
	Geometry   geometry
}

type geometry struct {
	X float64
	Y float64
}


var data query

func Query(u string, p bytes.Buffer) query {
	//[]features {

	response, err := http.Post(u, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(p.String())))
	if err != nil {
		fmt.Printf("%s", err)

	} else {
		defer response.Body.Close()
		c, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(c, &data)

	}
	return data //.Features
}

func main() {

  //This allows for appending more attributes because of not knowing them all ahead of time... Juts CYA
	var body bytes.Buffer
	body.WriteString("where=1=1&outFields=*&f=json&outSR=4326")

	//Pass parameters in URL. MUST HAVE -----> f=json
	d := Query("http://coagisweb.cabq.gov/arcgis/rest/services/public/PublicArt/MapServer/0/query", body)

	//fmt.Println(d.Features[5].Geometry)
	fmt.Println(d.DisplayFieldName)
	for i := 0; i < len(d.Features); i++ {
		fmt.Println(d.Features[i].Attributes["TITLE"])
		fmt.Println(d.Features[i].Geometry.Y)
		fmt.Println(d.Features[i].Geometry.X)

	}
}

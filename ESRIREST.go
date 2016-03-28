package github.com/MunicipalDevelopment/ESRI-GO/esrirest

import (
  "fmt" 
  "net/http"
  "io/ioutil"
  "encoding/json"
  "bytes"  

)


//in ESRI REST, you want the features[] - this is where results are stored (other info is metadata).
  type query struct{
  Features []features
  
  }

//Features has another [] of attributes. You can either hardoce all the values like "OBJECTID string" or if you want more generic, you can map it. then below you can iterate.

type features struct{
  Attributes map[string]interface{}
  Geometry geometry
}

type geometry struct{
  X float64
  Y float64
  
}

//This allows for appending more attributes because of not knowing them all ahead of time... Juts CYA


var data query

func Query(u string, p bytes.Buffer) []features {



        response, err := http.Post(u,"application/x-www-form-urlencoded",bytes.NewBuffer([]byte(p.String())))
         if err != nil {
             fmt.Printf("%s", err)

         } else {
             defer response.Body.Close()
             c, _ := ioutil.ReadAll(response.Body)
             json.Unmarshal(c,&data)





          }
  return data.Features
}

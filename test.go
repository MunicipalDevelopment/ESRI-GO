package main

import (
  "fmt" 
  "net/http"
  "io/ioutil"
  "encoding/json"
  "bytes"  

)


func main(){
var body bytes.Buffer
body.WriteString("where=1=1&outFields=*&f=json&outSR=4326")
 //Pass query parameters in URL. MUST HAVE -----> f=json
d :=Query("http://coagisweb.cabq.gov/arcgis/rest/services/public/PublicArt/MapServer/0/query",body)
//fmt.Println(d[5].Geometry)
 
      for i:=0; i<len(d); i++{
           fmt.Println(d[i].Attributes["TITLE"])
	//fmt.Println(d[i].Geometry.Y)
      
}
}

package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)
// parsing the data adnd unmarshalling the data from json 
func ParseBody(r *http.Request, x interface{}){
	//reading the body
	if body,err:=ioutil.ReadAll(r.Body);err == nil{
		//unmahsalling the data and converting it into slice of byte
		if err:= json.Unmarshal([]byte(body),x);err !=nil{
			return 
		}
	}
}
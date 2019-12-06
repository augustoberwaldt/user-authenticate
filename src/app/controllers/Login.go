package controllers

import (
	"../dto"
	"encoding/json"
	"net/http"
)



func  Authenticate(reponse http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var payload dto.Request_authenticate
	err := decoder.Decode(&payload)

	if err != nil {
		panic(err)
	}

	res := dto.Response {
		Code: "121243234",
		Msg: "sascascasc",
	}

	reponse.Header().Set("Content-Type", "application/json")

	json.NewEncoder(reponse).Encode(res)
}

package main

import (
	"encoding/json"
	"net/http"
)

//User defines model for storing account details in database
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", echoHandler)

	http.ListenAndServe(":8080", mux)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	result := Result{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	//err := json.NewDecoder(r.Body).Decode(&user)
	//if err != nil{
	//	panic(err)
	//}

	//Set CreatedAt field on user to current local time
	//user.CreatedAt = time.Now().Local()

	result.Code = 0
	result.Msg = "success"
	result.Data = r.URL.Path

	//Marshal or convert user object back to json and write to response
	resultJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(resultJson)
}

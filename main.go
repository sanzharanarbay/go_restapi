package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Name    string
}

type Params struct {
	ID   int
	User string
}

func test(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		response:= Response{"The method is not post"}
		json,error := json.Marshal(response)
		if error != nil {
			http.Error(w, error.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	p := &Params{}
	err = json.Unmarshal(body, p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "content-type %#v\n",
		r.Header.Get("Content-Type"))
	fmt.Fprintf(w, "params %#v\n", p)

}

func main() {
	http.HandleFunc("/test", test)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

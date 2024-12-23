package application

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"github.com/MaksaNeNegr/calc_go/pkg/rpn"
)
///
type Application struct {
}

func New() *Application {
	return &Application{}
}

type Request struct {
	Expression string `json: "expression"`
}

type Response struct {
	Res string `json:"result"`
}

type Err_Response struct {
	Error_ string `json:"error"`
}

type error interface {
    Error() string
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ожидается post-зарос", 500)
		return // единственный случай, когда возращяется код 500
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&req)
	res_, err := rpn.Calc(req.Expression)
	w.Header().Set("Content-Type", "application/json")
	if err == nil{
		w.WriteHeader(200)
		var res Response
		res.Res = res_
		json.NewEncoder(w).Encode(res)
	} else {
		w.WriteHeader(422)
		var err_ Err_Response
		err_.Error_ = err.Error()
		json.NewEncoder(w).Encode(err_)
	}
	// fmt.Printf(res_)
	// fmt.Println(req.Expression)

	// w.Header().Set("Content-Type", "application/json")

	
}

func (a *Application) Run() { 
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	http.ListenAndServe(":8080", nil)
}

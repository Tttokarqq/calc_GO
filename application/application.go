package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/MaksaNeNegr/calc_go/pkg/rpn"
)


type express struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type ListExpress struct {
	Expressions []express `json:"expressions"` //
}

var(
	id = 0
	List1 = ListExpress{[]express{}}
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

type Request struct {
	Expression string `json: "expression"`
}

type Response1 struct {
	Res string `json:"id"`
}

type Err_Response struct {
	Error_ string `json:"error"`
}

type error interface {
    Error() string
}

func retErr(w http.ResponseWriter, code int, err error){
	w.WriteHeader(code)
	var err_ Err_Response
	err_.Error_ = err.Error()
	json.NewEncoder(w).Encode(err_)
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	retErr(w, 500, rpn.Err_no_post)
	// 	return // единственный случай, когда возращяется код 500
	// }

	var req Request
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&req)
	r_, err := rpn.Calc2(req.Expression)
	res_ := strconv.Itoa(id)
	id += 1
	w.Header().Set("Content-Type", "application/json")
	if err == nil{
		ex1 := express{res_, "wait", r_}
		List1.Expressions = append(List1.Expressions, ex1)
		w.WriteHeader(200)
		var res Response1
		res.Res = res_
		json.NewEncoder(w).Encode(res)
	} else {
		retErr(w, 422, err)
	}
	// fmt.Printf(res_)
	// fmt.Println(req.Expression)

	// w.Header().Set("Content-Type", "application/json")
}

func accuracy_(w http.ResponseWriter, r *http.Request) {
	acc := r.URL.Query().Get("accuracy")
	acc_, err := strconv.Atoi(acc)
	if err != nil{
		retErr(w, 405, rpn.Err_acc)
	} else if acc_ >= 65 || acc_ < 0{
		retErr(w, 405, rpn.Err_acc)
	} else {
		// acc, _ = strconv.Itoa(acc_ + 2) // почему то всегда округляется на 2 цифры меньше???
		rpn.ChangeTochonst(strconv.Itoa(acc_ + 2))
	}
}

func showEx(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)                      
	// fmt.Println(rpn.Tasks)
	
	json.NewEncoder(w).Encode(List1)
}



func getWork(w http.ResponseWriter, r *http.Request) {
	fmt.Println(rpn.Tasks)
    // for _, task := range List1 {
    //     if _, err1 := strconv.ParseFloat(task.Arg1, 64); err1 == nil {
    //         if _, err2 := strconv.ParseFloat(task.Arg2, 64); err2 == nil {
    //             response := map[string]interface{}{
    //                 "task": task,
    //             }

    //             w.Header().Set("Content-Type", "application/json")
    //             w.WriteHeader(http.StatusOK)

    //             if err := json.NewEncoder(w).Encode(response); err != nil {
    //                 http.Error(w, err.Error(), http.StatusInternalServerError)
    //             }
    //             return
    //         }
    //     }
    // }
}

func (a *Application) Run() { 
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	// http.HandleFunc("/api/v1/calculate/acc", accuracy_)
	http.HandleFunc("/api/v1/calculate/acc", accuracy_)
	http.HandleFunc("/api/v1/expressions", showEx)
	http.HandleFunc("/internal/task", getWork)
	http.ListenAndServe(":8080", nil)
}
// curl http://localhost:8080/api/v1/calculate/acc?accuracy=2

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ErrType interface {
	~bool | ~string
}

func main() {
	router := httprouter.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.HandlerFunc(http.MethodGet, "/api/classify-number", classifyNumber)
	log.Printf("Server connected on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	log.Fatal(err)

}

type ErrMessage struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

type SucessResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

type Parity int // a fixated type on integer

// classify number api router
func classifyNumber(w http.ResponseWriter, r *http.Request) {
	que := r.URL.Query().Get("number")
	param, err := strconv.Atoi(que)
	if err != nil {
		errMess := ErrMessage{Number: que, Error: true}
		bt, err := json.Marshal(errMess)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(bt))
		return
	}
	par := Parity(param)
	data, err := sendRequest(par)

	if err != nil {
		errMess := ErrMessage{Number: que, Error: true}
		bt, err := json.Marshal(errMess)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(bt))
		return
	}

	// collect data
	resp := struct {
		Number     Parity   `json:"number"`
		IsPrime    bool     `json:"is_prime"`
		IsPerfect  bool     `json:"is_perfect"`
		Properties []string `json:"properties"`
		DigitSum   int      `json:"digit_sum"`
		FunFact    string   `json:"fun_fact"`
	}{
		Number:     par,
		IsPrime:    par.checkIsPrime(),
		IsPerfect:  par.isPerfectNumber(),
		Properties: numberProperties(par),
		DigitSum:   par.calcSumOfNumbers(),
		FunFact:    data,
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jsonData))
}

// send request to numbers api
func sendRequest(qParam Parity) (string, error) {

	params := url.Values{}

	params.Add("", fmt.Sprint(qParam))
	fullUrl := fmt.Sprintf("http://numbersapi.com/%d/math", qParam)
	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)

	if err != nil {
		return "", err
	}
	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	stringData := string(data)
	return stringData, nil
}

// checks if a number is either even or odd
func numberProperties(num Parity) []string {

	prop := make([]string, 0)

	if num.checkNumberIsArmStrong() == true {
		prop = append(prop, "armstrong")
	}
	prop = append(prop, num.checkParity())

	return prop
}

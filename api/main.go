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
	param, err := strconv.Atoi(r.URL.Query().Get("number"))
	// throw a bad response error
	if err != nil {
		errMess := ErrMessage{Number: "alphabet", Error: true}
		bt, err := json.Marshal(errMess)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(bt))
		return
	}
	a, err := sendRequest(param)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
	w.Write([]byte(fmt.Sprintf("%d", param)))
}

// send user request
func sendRequest(qParam int) (*http.Response, error) {

	params := url.Values{}

	params.Add("", fmt.Sprint(qParam))
	fullUrl := fmt.Sprintf("http://numbersapi.com/%d/math?json", qParam)
	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)

	if err != nil {
		return nil, err
	}
	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(data))
	return resp, nil
}

// checks if a number is either even or odd

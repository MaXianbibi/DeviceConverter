package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	apiKey = "cur_live_8qyVSM7iTLtm2rJzEus96HrRk8xt5rB1vlTzx4zx"
)

import (
	"encoding/json"
)

type Currency struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

type Data struct {
	EUR Currency `json:"EUR"`
}

type Response struct {
	Meta struct {
		LastUpdatedAt string `json:"last_updated_at"`
	} `json:"meta"`
	Data Data `json:"data"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Hello, please enter a number to convert: ")
	number, _ := reader.ReadString('\n')
	number = number[:len(number)-1]

	numberFloat, err := strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}

	fmt.Print("First Currency: ")
	firstCurrency, _ := reader.ReadString('\n')
	firstCurrency = firstCurrency[:len(firstCurrency)-1]

	fmt.Print("Second Currency: ")
	secondCurrency, _ := reader.ReadString('\n')
	secondCurrency = secondCurrency[:len(secondCurrency)-1]

	url := fmt.Sprintf("https://api.currencyapi.com/v3/latest?&base_currency=%s&currencies=%s", firstCurrency, secondCurrency)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("apikey", apiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.Data.EUR.Value) // Affiche la valeur de EUR.

	result := data.Data.EUR.Value * numberFloat
	fmt.Println("Result: ", result)
}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello, please enter a number to convert : ")
	number, _ := reader.ReadString('\n')
	number = number[:len(number)-1]

	numberFloat, err := strconv.ParseFloat(number, 64)
	if (err != nil) { 
		fmt.Println("Please enter a valid number")
		return
	}


	fmt.Print("First Currency : ")
	firstCurrency, _ := reader.ReadString('\n')
	firstCurrency = firstCurrency[:len(firstCurrency)-1]

	fmt.Print("Second Currency : ")
	secondCurrency, _ := reader.ReadString('\n')
	secondCurrency = secondCurrency[:len(secondCurrency)-1]

	url := "https://api.currencyapi.com/v3/latest?&base_currency=" + firstCurrency + "&currencies=" + secondCurrency

		fmt.Println(url)
	var method string = "GET"
	// https://api.currencyapi.com/v3/latest?apikey=cur_live_8qyVSM7iTLtm2rJzEus96HrRk8xt5rB1vlTzx4zx&currencies=EUR%2CUSD

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("apikey", "cur_live_8qyVSM7iTLtm2rJzEus96HrRk8xt5rB1vlTzx4zx")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var data Response
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(data.Data.EUR.Value) // Affiche la valeur de EUR.


	resultat := data.Data.EUR.Value * numberFloat
	fmt.Println("Resultat : ", resultat)
}

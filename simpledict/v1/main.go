package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) error {

	url := "https://api.interpreter.caiyunai.com/v1/dict"

	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json.Marshal %w", err)
	}
	payload := bytes.NewReader(buf)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return fmt.Errorf("http.NewRequest %w", err)
	}

	req.Header.Add("user-agent", "vscode-restclient")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("x-authorization", "token:qgemv4jr1y38jyq6vhvi")

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http.DefaultClient.Do %w", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll %w", err)
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("bad body: %s", body)
	}
	// fmt.Println(res)
	// fmt.Println(string(body))
	var resp DictResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return fmt.Errorf("json.Unmarshal %w", err)
	}
	// fmt.Println(resp)
	fmt.Println(word, "UK:", resp.Dictionary.Prons.En, "US:", resp.Dictionary.Prons.EnUs)
	for _, item := range resp.Dictionary.Explanations {
		fmt.Println(item)
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	err := query(word)
	if err != nil {
		fmt.Printf("query error: %+v\n", err)
	}
}

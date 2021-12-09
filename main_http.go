package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func MainHttpPostJsonWithHeader() {
	// 初始化client
	var client = &http.Client{}

	// 请求地址
	var url = "http://restapi3.apiary.io/notes"

	// 请求数据
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)

	// 初始化一个请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Print(err)
	}

	// 设置请求头
	req.Header.Set("target-service-name", "sadasdasda")

	// 发出请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Print(body, err)
}

func MainHttpGet() {
	name := "John Doe"
	occupation := "gardener"
	params := "name=" + url.QueryEscape(name) + "&" +
		"occupation=" + url.QueryEscape(occupation)
	path := fmt.Sprintf("https://httpbin.org/get?%s", params)

	resp, err := http.Get(path)
	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Print(body, err)
}

func MainHttpPostForm() {
	data := url.Values{
		"name":       {"John Doe"},
		"occupation": {"gardener"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", data)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])
}

func MainHttpPostJson() {
	values := map[string]string{"name": "John Doe", "occupation": "gardener"}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}
